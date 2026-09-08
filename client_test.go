package easylabsdk

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"connectrpc.com/connect"

	agentv1 "github.com/abcp-sdk/agent-proto/agent/v1"
	"github.com/abcp-sdk/agent-proto/agent/v1/agentv1connect"
	easylabv1 "github.com/easylab-platform/easylab-proto/easylab/v1"
	"github.com/easylab-platform/easylab-proto/easylab/v1/easylabv1connect"
)

// stubLab implements the Lab surface pieces the tests exercise.
type stubLab struct {
	easylabv1connect.UnimplementedLabServiceHandler
	ensureCalls atomic.Int32
}

func (s *stubLab) Health(
	ctx context.Context, req *connect.Request[easylabv1.HealthRequest],
) (*connect.Response[easylabv1.HealthResponse], error) {
	return connect.NewResponse(&easylabv1.HealthResponse{Ok: true}), nil
}

func (s *stubLab) EnsureRepo(
	ctx context.Context, req *connect.Request[easylabv1.EnsureRepoRequest],
) (*connect.Response[easylabv1.EnsureRepoResponse], error) {
	n := s.ensureCalls.Add(1)
	if n > 1 {
		return nil, connect.NewError(connect.CodeAlreadyExists, errStr("repo exists"))
	}
	return connect.NewResponse(&easylabv1.EnsureRepoResponse{Ok: true}), nil
}

func (s *stubLab) ListRepos(
	ctx context.Context, req *connect.Request[easylabv1.ListReposRequest],
) (*connect.Response[easylabv1.ListReposResponse], error) {
	return connect.NewResponse(&easylabv1.ListReposResponse{
		Repos: []*easylabv1.RepoInfo{
			{Namespace: "acme", Name: "api", DefaultBranch: "main", Sha: "deadbeef"},
		},
	}), nil
}

// stubAgentForward mirrors the gateway's /agent.v1.* forward target.
type stubAgentForward struct {
	agentv1connect.UnimplementedAgentServiceHandler
}

func (s *stubAgentForward) ListSessions(
	ctx context.Context, req *connect.Request[agentv1.ListSessionsRequest],
) (*connect.Response[agentv1.ListSessionsResponse], error) {
	return connect.NewResponse(&agentv1.ListSessionsResponse{
		Sessions: []*agentv1.Session{{Name: "forwarded"}},
	}), nil
}

type recorder struct {
	proto atomic.Value
	auth  atomic.Value
}

func newGatewayTestServer(t *testing.T) (*httptest.Server, *recorder, *stubLab) {
	t.Helper()
	rec := &recorder{}
	lab := &stubLab{}
	mux := http.NewServeMux()
	mux.Handle(easylabv1connect.NewLabServiceHandler(lab))
	mux.Handle(agentv1connect.NewAgentServiceHandler(&stubAgentForward{}))
	srv := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec.proto.Store(r.Proto)
		rec.auth.Store(r.Header.Get("Authorization"))
		mux.ServeHTTP(w, r)
	}))
	srv.Start()
	t.Cleanup(srv.Close)
	return srv, rec, lab
}

// TestGatewayRoundTrip drives Lab helpers + the Agent forward surface
// through the SDK against a stub gateway over HTTP/1.1.
func TestGatewayRoundTrip(t *testing.T) {
	srv, rec, lab := newGatewayTestServer(t)
	client := New(srv.URL, "tok-9")

	if err := client.EnsureRepo(context.Background(), "acme", "api"); err != nil {
		t.Fatalf("EnsureRepo: %v", err)
	}
	if err := client.EnsureRepo(context.Background(), "acme", "api"); err != nil {
		t.Fatalf("EnsureRepo must be idempotent, got %v", err)
	}
	if lab.ensureCalls.Load() != 2 {
		t.Fatalf("ensureCalls = %d, want 2", lab.ensureCalls.Load())
	}

	repos, err := client.ListRepos(context.Background())
	if err != nil {
		t.Fatalf("ListRepos: %v", err)
	}
	if len(repos) != 1 || repos[0].Namespace != "acme" || repos[0].Name != "api" {
		t.Fatalf("repos = %+v", repos)
	}

	// The agent surface rides the same gateway (forwarded, not direct).
	sessions, err := client.Agent.ListSessions(
		context.Background(), connect.NewRequest(&agentv1.ListSessionsRequest{}),
	)
	if err != nil {
		t.Fatalf("Agent.ListSessions via gateway: %v", err)
	}
	if len(sessions.Msg.GetSessions()) != 1 || sessions.Msg.GetSessions()[0].GetName() != "forwarded" {
		t.Fatalf("sessions = %v", sessions.Msg.GetSessions())
	}

	if got := rec.proto.Load(); got != "HTTP/1.1" {
		t.Fatalf("protocol = %v, want HTTP/1.1 (default client)", got)
	}
	if got := rec.auth.Load(); got != "Bearer tok-9" {
		t.Fatalf("authorization = %v, want Bearer tok-9", got)
	}
}

// TestWithHTTPClientOverride proves the option wiring via a client whose
// transport records the dial.
func TestWithHTTPClientOverride(t *testing.T) {
	srv, _, _ := newGatewayTestServer(t)

	dialed := atomic.Bool{}
	custom := &http.Client{Transport: recordingTransport{&dialed}}
	client := New(srv.URL, "", WithHTTPClient(custom))

	if _, err := client.ListRepos(context.Background()); err != nil {
		t.Fatalf("ListRepos: %v", err)
	}
	if !dialed.Load() {
		t.Fatal("custom http.Client was not used")
	}
}

type recordingTransport struct{ dialed *atomic.Bool }

func (rt recordingTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	rt.dialed.Store(true)
	return http.DefaultTransport.RoundTrip(r)
}

func errStr(s string) error { return &strErr{s} }

type strErr struct{ msg string }

func (e *strErr) Error() string { return e.msg }
