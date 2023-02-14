package mec

const (
	// esHostNameField is the field name of host name in the document
	esHostNameField = "host.hostname"
	// esCpuUsageField is the field name of cpu usage in the document
	esCpuUsageField = "host.cpu.usage"
	// esMemUsageField is the field name of mem usage in the document
	esMemUsageField = "system.memory.actual.used.pct"
)

type ElasticsearchMetricsClient struct {
	address string
}

func NewElasticsearchMetricsClient(address string) (*ElasticsearchMetricsClient, error) {
	return &ElasticsearchMetricsClient{address: address}, nil
}

func (p *ElasticsearchMetricsClient) NodeMetricsAvg(nodeName string, period string) (*NodeMetrics, error) {
	nodeMetrics := &NodeMetrics{}
	return nodeMetrics, nil
}
