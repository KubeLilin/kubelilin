package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sgr/pkg/pipeline"
	"testing"
)

func TestJenkinsJob(t *testing.T) {
	builder := pipeline.NewBuilder()
	builder.UseJenkins("http://152.136.141.235:32001", "jenkins", "11e681bb454a36a9ce0e0a6fd030d059a9").
		UseKubernetes("sgr-ci", "golang:1.16.5")

	pipeline, _ := builder.Build()

	ping, err := pipeline.Ping()
	if err != nil {
		return
	}
	fmt.Printf("jenkins version  %s", ping)

	assert.Equal(t, ping != "", true)

	pipeline.SwitchJobName("sample-pipeline-test")
	job, _ := pipeline.GetJobInfo(11)

	assert.Equal(t, job != nil, true)

	logs, err := pipeline.GetJobLogs(11)

	assert.Equal(t, logs != "", true)

}
