package metrics

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"time"
)

type PrometheusMetrics struct {
	DataSource string `json:"dataSource"`
	clientApi  v1.API
}

func NewPrometheusMetrics(dataSource string) *PrometheusMetrics {
	client, err := api.NewClient(api.Config{
		Address: dataSource,
	})
	if err != nil {
		panic(err)
	}
	return &PrometheusMetrics{
		clientApi:  v1.NewAPI(client),
		DataSource: dataSource,
	}
}

func (metrics *PrometheusMetrics) Query(query string, startTime time.Time, endTime time.Time) (string, error) {
	fmt.Println(query)
	result, _, err := metrics.clientApi.QueryRange(context.Background(), query,
		v1.Range{
			Start: startTime,
			End:   endTime,
			Step:  time.Minute,
		})
	if err != nil {
		return "", err
	}
	data, err := json.Marshal(result)
	return string(data), err
}
