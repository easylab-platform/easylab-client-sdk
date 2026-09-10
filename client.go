// Package easylabsdk re-exports the buf-generated EasyLab easylab.v1 (and
// worker.v1) message types and Connect client constructors.
//
// Nothing here is hand-written: for the clients themselves import the
// generated packages directly:
//
//	easylabv1 "github.com/easylab-platform/easylab-proto/easylab/v1"
//	"github.com/easylab-platform/easylab-proto/easylab/v1/easylabv1connect"
//
//	client := easylabv1connect.NewLabServiceClient(httpClient, baseURL, opts...)
//
// This package exists only so consumers have the common names under one
// import.
package easylabsdk

import (
	easylabv1 "github.com/easylab-platform/easylab-proto/easylab/v1"
	"github.com/easylab-platform/easylab-proto/easylab/v1/easylabv1connect"
	workerv1 "github.com/easylab-platform/easylab-proto/worker/v1"
)

type (
	// Message aliases (easylab.v1).
	RepoInfo     = easylabv1.RepoInfo
	Ok           = easylabv1.Ok
	FileEntry    = easylabv1.FileEntry
	DiffFile     = easylabv1.DiffFile
	CommitInfo   = easylabv1.CommitInfo
	TagInfo      = easylabv1.TagInfo
	BranchInfo   = easylabv1.BranchInfo
	RevisionInfo = easylabv1.RevisionInfo
	ServiceInfo  = easylabv1.ServiceInfo
	TaskEntry    = easylabv1.TaskEntry
	SandboxInfo  = easylabv1.SandboxInfo
	Workflow     = easylabv1.Workflow
	JobDef       = easylabv1.JobDef
	Step         = easylabv1.Step
	Produce      = easylabv1.Produce
	Trigger      = easylabv1.Trigger
	PortSpec     = easylabv1.PortSpec

	// Request/response aliases (easylab.v1).
	ListReposRequest    = easylabv1.ListReposRequest
	EnsureRepoRequest   = easylabv1.EnsureRepoRequest
	EnsureOrgRequest    = easylabv1.EnsureOrgRequest
	BranchesRequest     = easylabv1.BranchesRequest
	RevisionsRequest    = easylabv1.RevisionsRequest
	CreateBranchRequest = easylabv1.CreateBranchRequest
	DeleteBranchRequest = easylabv1.DeleteBranchRequest
	ReadBlobRequest     = easylabv1.ReadBlobRequest
	WriteBlobRequest    = easylabv1.WriteBlobRequest
	LogRequest          = easylabv1.LogRequest
	DiffRequest         = easylabv1.DiffRequest
	SearchRequest       = easylabv1.SearchRequest
	GraphRequest        = easylabv1.GraphRequest
	CompareRequest      = easylabv1.CompareRequest
	RebaseRequest       = easylabv1.RebaseRequest
	TreeRequest         = easylabv1.TreeRequest
	BlameRequest        = easylabv1.BlameRequest
	SetMirrorRequest    = easylabv1.SetMirrorRequest
	ListPackageTypesRequest = easylabv1.ListPackageTypesRequest
	ListServicesRequest = easylabv1.ListServicesRequest
	GetServiceRequest   = easylabv1.GetServiceRequest
	GetTaskRequest      = easylabv1.GetTaskRequest
	TaskLogRequest      = easylabv1.TaskLogRequest
	LaunchServiceRequest = easylabv1.LaunchServiceRequest
	DeleteServiceRequest = easylabv1.DeleteServiceRequest
	ScaleServiceRequest  = easylabv1.ScaleServiceRequest
	SyncRequest          = easylabv1.SyncRequest
	CreateWorkflowRequest = easylabv1.CreateWorkflowRequest
	TriggerRunRequest    = easylabv1.TriggerRunRequest

	// Sandbox request aliases (easylab.v1 wrappers over worker.v1).
	EnsureSandboxImageRequest = easylabv1.EnsureSandboxImageRequest
	LaunchSandboxRequest      = easylabv1.LaunchSandboxRequest
	GetSandboxRequest         = easylabv1.GetSandboxRequest
	DeleteSandboxRequest      = easylabv1.DeleteSandboxRequest
	ListSandboxesRequest      = easylabv1.ListSandboxesRequest
	ExecuteRequest            = easylabv1.ExecuteRequest
	ListJobsRequest           = easylabv1.ListJobsRequest
	JobOutputRequest          = easylabv1.JobOutputRequest
	WatchJobRequest           = easylabv1.WatchJobRequest
	JobWaitRequest            = easylabv1.JobWaitRequest
	JobStdinRequest           = easylabv1.JobStdinRequest
	JobKillRequest            = easylabv1.JobKillRequest
	FileReadRequest           = easylabv1.FileReadRequest
	FileWriteRequest          = easylabv1.FileWriteRequest
	FileListRequest           = easylabv1.FileListRequest
	SyncWorkspaceRequest      = easylabv1.SyncWorkspaceRequest

	// worker.v1 payload aliases (the req field of the sandbox RPCs).
	WorkerExecuteRequest   = workerv1.ExecuteRequest
	WorkerJobOutputRequest = workerv1.JobOutputRequest
	WorkerJobWaitRequest   = workerv1.JobWaitRequest
	WorkerJobStdinRequest  = workerv1.JobStdinRequest
	WorkerJobKillRequest   = workerv1.JobKillRequest
	WorkerFileReadRequest  = workerv1.FileReadRequest
	WorkerFileWriteRequest = workerv1.FileWriteRequest
	WorkerFileListRequest  = workerv1.FileListRequest
	WorkerFileEntry        = workerv1.FileEntry
)

var (
	// Service client constructors (generated).
	NewLabServiceClient      = easylabv1connect.NewLabServiceClient
	NewOpsServiceClient      = easylabv1connect.NewOpsServiceClient
	NewRegistryServiceClient = easylabv1connect.NewRegistryServiceClient
	NewSandboxServiceClient  = easylabv1connect.NewSandboxServiceClient
	NewWorkflowServiceClient = easylabv1connect.NewWorkflowServiceClient
)
