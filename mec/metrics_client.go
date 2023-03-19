package mec

import (
	"context"
	"errors"
)

type MetricsClient interface {
	NodeMetricsAvg(ctx context.Context, nodeName string, period string) (*NodeMetrics, error)
}

func NewMetricsClient(metricsConf map[string]string) (MetricsClient, error) {
	address := metricsConf["address"]
	if len(address) == 0 {
		return nil, errors.New("metrics address is empty")
	}
	metricsType := metricsConf["type"]
	if metricsType == "elasticsearch" {
		return NewElasticsearchMetricsClient(address, metricsConf)
	}
	return NewPrometheusMetricsClient(address, metricsConf)
}
