package mec

type MetricsClient interface {
	NodeMetricsAvg(nodeName string, period string) (*NodeMetrics, error)
}
