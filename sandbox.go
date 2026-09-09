package easylabsdk

import (
	"context"

	"connectrpc.com/connect"
	easylabv1 "github.com/easylab-platform/easylab-proto/easylab/v1"
	workerv1 "github.com/easylab-platform/easylab-proto/worker/v1"
)

// sandbox is the SandboxService surface with ergonomic wrappers over the
// generated Connect client. Every call routes a sandbox name + workerv1.v1
// payload. The workerv1.v1 message types are reused as-is.
func (c *Client) Sandbox() sandboxClient {
	return sandboxClient{c: c}
}

type sandboxClient struct {
	c *Client
}

// EnsureSandboxImage caches a derived image (base + injected worker).
func (s sandboxClient) EnsureSandboxImage(ctx context.Context, baseImage string) (*connect.Response[easylabv1.EnsureSandboxImageResponse], error) {
	return s.c.SandboxService.EnsureSandboxImage(ctx, connect.NewRequest(&easylabv1.EnsureSandboxImageRequest{BaseImage: baseImage}))
}

// LaunchSandbox launches a worker-backed sandbox from an arbitrary base image.
func (s sandboxClient) LaunchSandbox(ctx context.Context, req *easylabv1.LaunchSandboxRequest) (*connect.Response[easylabv1.LaunchSandboxResponse], error) {
	return s.c.SandboxService.LaunchSandbox(ctx, connect.NewRequest(req))
}

// DeleteSandbox removes the sandbox container + registration.
func (s sandboxClient) DeleteSandbox(ctx context.Context, name string) (*connect.Response[easylabv1.DeleteSandboxResponse], error) {
	return s.c.SandboxService.DeleteSandbox(ctx, connect.NewRequest(&easylabv1.DeleteSandboxRequest{Name: name}))
}

// GetSandbox merges registry row + live status + worker stats.
func (s sandboxClient) GetSandbox(ctx context.Context, name string) (*connect.Response[easylabv1.GetSandboxResponse], error) {
	return s.c.SandboxService.GetSandbox(ctx, connect.NewRequest(&easylabv1.GetSandboxRequest{Name: name}))
}

func (s sandboxClient) ListSandboxes(ctx context.Context) (*connect.Response[easylabv1.ListSandboxesResponse], error) {
	return s.c.SandboxService.ListSandboxes(ctx, connect.NewRequest(&easylabv1.ListSandboxesRequest{}))
}

// Execute runs a command in the sandbox, returning a job id.
func (s sandboxClient) Execute(ctx context.Context, sandbox string, req *workerv1.ExecuteRequest) (*connect.Response[workerv1.ExecuteResponse], error) {
	return s.c.SandboxService.Execute(ctx, connect.NewRequest(&easylabv1.ExecuteRequest{Sandbox: sandbox, Req: req}))
}

// ListJobs lists jobs in the sandbox.
func (s sandboxClient) ListJobs(ctx context.Context, sandbox string) (*connect.Response[workerv1.ListJobsResponse], error) {
	return s.c.SandboxService.ListJobs(ctx, connect.NewRequest(&easylabv1.ListJobsRequest{Sandbox: sandbox}))
}

// JobOutput paginates job history (start may be negative; stream all|stdout|stderr).
func (s sandboxClient) JobOutput(ctx context.Context, sandbox string, req *workerv1.JobOutputRequest) (*connect.Response[workerv1.JobOutputResponse], error) {
	return s.c.SandboxService.JobOutput(ctx, connect.NewRequest(&easylabv1.JobOutputRequest{Sandbox: sandbox, Req: req}))
}

// JobWait blocks until a job completes or times out.
func (s sandboxClient) JobWait(ctx context.Context, sandbox string, req *workerv1.JobWaitRequest) (*connect.Response[workerv1.JobWaitResponse], error) {
	return s.c.SandboxService.JobWait(ctx, connect.NewRequest(&easylabv1.JobWaitRequest{Sandbox: sandbox, Req: req}))
}

// JobKill kills a job in the sandbox.
func (s sandboxClient) JobKill(ctx context.Context, sandbox string, req *workerv1.JobKillRequest) (*connect.Response[workerv1.JobKillResponse], error) {
	return s.c.SandboxService.JobKill(ctx, connect.NewRequest(&easylabv1.JobKillRequest{Sandbox: sandbox, Req: req}))
}

// JobStdin writes to a job's stdin (closeAfter ends the pipe).
func (s sandboxClient) JobStdin(ctx context.Context, sandbox string, req *workerv1.JobStdinRequest) (*connect.Response[workerv1.JobStdinResponse], error) {
	return s.c.SandboxService.JobStdin(ctx, connect.NewRequest(&easylabv1.JobStdinRequest{Sandbox: sandbox, Req: req}))
}

// FileRead reads a sandbox file (bytes).
func (s sandboxClient) FileRead(ctx context.Context, sandbox string, req *workerv1.FileReadRequest) (*connect.Response[workerv1.FileReadResponse], error) {
	return s.c.SandboxService.FileRead(ctx, connect.NewRequest(&easylabv1.FileReadRequest{Sandbox: sandbox, Req: req}))
}

// FileWrite writes a sandbox file (bytes).
func (s sandboxClient) FileWrite(ctx context.Context, sandbox string, req *workerv1.FileWriteRequest) (*connect.Response[workerv1.FileWriteResponse], error) {
	return s.c.SandboxService.FileWrite(ctx, connect.NewRequest(&easylabv1.FileWriteRequest{Sandbox: sandbox, Req: req}))
}

// FileList lists a sandbox path.
func (s sandboxClient) FileList(ctx context.Context, sandbox string, req *workerv1.FileListRequest) (*connect.Response[workerv1.FileListResponse], error) {
	return s.c.SandboxService.FileList(ctx, connect.NewRequest(&easylabv1.FileListRequest{Sandbox: sandbox, Req: req}))
}


// SyncWorkspace pushes the repo tree at rev into the sandbox and records
// rev + worker boot id in the registry (idempotent; easylab skips when the
// sandbox is already synced at this rev/boot).
func (s sandboxClient) SyncWorkspace(ctx context.Context, req *easylabv1.SyncWorkspaceRequest) (*connect.Response[easylabv1.SyncWorkspaceResponse], error) {
	return s.c.SandboxService.SyncWorkspace(ctx, connect.NewRequest(req))
}
