package pipeline

import (
	"time"

	httpclient "github.com/isbrick/http-client"
)

// Driver ...
type Driver int

const (
	DriverJenkins Driver = iota + 1
)

func (d Driver) String() (s string) {
	switch d {
	case DriverJenkins:
		return "jenkins"
	default:
		return "unknown"
	}
}

// Pipeline ..
type Pipeline interface {
	Ping() (string, error)
	Abort(jobName string, RunID int64) error
	GetJobInfo(jobName string, runID int64) (*JobInfo, error)
	GetJobLogs(jobName string, runID int64) (string, error)
	RunJob(jobName string) (int64, error)
	RunJobWithParameters(jobName string, branch string) (int64, error)
	SaveJob(jobName string, processor FlowProcessor) error
	//SwitchJobName(jobName string)
	//Build() (int64, error)
	//SetWorkFlow(processor FlowProcessor)
}

// HTTPClient defined http native client
var (
	timeout    = 15000 * time.Millisecond
	HTTPClient = httpclient.NewHClient(httpclient.WithHTTPTimeout(timeout))
)

// JobInfo ...
type JobInfo struct {
	Artifacts         []interface{} `json:"deliverables"`
	Building          bool          `json:"building"`
	Description       interface{}   `json:"description"`
	DisplayName       string        `json:"displayName"`
	Duration          int           `json:"duration"`
	EstimatedDuration int           `json:"estimatedDuration"`
	Executor          interface{}   `json:"executor"`
	FullDisplayName   string        `json:"fullDisplayName"`
	ID                string        `json:"id"`
	Number            int           `json:"number"`
	QueueID           int           `json:"queueId"`
	Result            string        `json:"result"`
	Status            string        `json:"status"`
	StartTimeMillis   int64         `json:"startTimeMillis"`
	EndTimeMillis     int64         `json:"endTimeMillis"`
	DurationMillis    int           `json:"durationMillis"`
	Stages            []Stage       `json:"stages"`
}

// Stage job's stage
type Stage struct {
	ID                  string `json:"id"`
	Name                string `json:"name"`
	ExecNode            string `json:"execNode"`
	Status              string `json:"status"`
	StartTimeMillis     int64  `json:"startTimeMillis"`
	DurationMillis      int    `json:"durationMillis"`
	PauseDurationMillis int    `json:"pauseDurationMillis"`
}
