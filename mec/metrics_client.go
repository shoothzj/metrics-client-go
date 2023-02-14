package mec

import "context"

type MetricsClient interface {
	NodeMetricsAvg(ctx context.Context, nodeName string, period string) (*NodeMetrics, error)
}
