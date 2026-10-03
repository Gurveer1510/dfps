package core

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	JobID      uuid.UUID
	ID         uuid.UUID
	FilePath   string
	Start      int64
	Length     int64
	OutputPath string
	State      TaskState
	WorkerID   string
	ClaimedAt  *time.Time
	AttemptNo  int
}

type TaskState string

const (
	TaskPending    TaskState = "PENDING"
	TaskProcessing TaskState = "PROCESSING"
	TaskDone       TaskState = "DONE"
	TaskFailed     TaskState = "FAILED"
)

type Job struct {
	ID         uuid.UUID
	InputFiles []string
	OutputPath string
	JobStatus  JobStatus
}

type JobStatus string

const (
	JobPending       JobStatus = "PENDING"
	JobChunksRunning JobStatus = "CHUNKS_PROCESSING"
	JobChunksMerging JobStatus = "CHUNKS_MERGING"
	JobCompleted     JobStatus = "COMPLETED"
	JobFailed        JobStatus = "FAILED"
	JobCancelled     JobStatus = "CANCELLED"
)
