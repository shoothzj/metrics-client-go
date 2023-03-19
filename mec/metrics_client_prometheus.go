package mec

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/prometheus/client_golang/api"
	prometheusv1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	// promCpuUsageAvg record name of cpu average usage defined in prometheus rules
	promCpuUsageAvg = "cpu_usage_avg"
	// promMemUsageAvg record name of mem average usage defined in prometheus rules
	promMemUsageAvg = "mem_usage_avg"
)

type PrometheusMetricsClient struct {
	address string
	conf    map[string]string
}

func NewPrometheusMetricsClient(address string, conf map[string]string) (*PrometheusMetricsClient, error) {
	return &PrometheusMetricsClient{address: address, conf: conf}, nil
}

func (p *PrometheusMetricsClient) NodeMetricsAvg(ctx context.Context, nodeName string, period string) (*NodeMetrics, error) {
	var client api.Client
	var err error
	insecureSkipVerify := p.conf["tls.insecureSkipVerify"] == "true"
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: insecureSkipVerify,
		},
	}
	client, err = api.NewClient(api.Config{
		Address:      p.address,
		RoundTripper: tr,
	})
	if err != nil {
		return nil, err
	}
	v1api := prometheusv1.NewAPI(client)
	nodeMetrics := &NodeMetrics{}
	for _, metric := range []string{promCpuUsageAvg, promMemUsageAvg} {
		queryStr := fmt.Sprintf("%s_%s{instance=\"%s\"}", metric, period, nodeName)
		res, warnings, err := v1api.Query(ctx, queryStr, time.Now())
		if err != nil {
			return nil, err
		}
		if res == nil || res.String() == "" {
			return nil, errors.New("no data")
		}
		if len(warnings) > 0 {
			return nil, err
		}
		firstRowValVector := strings.Split(res.String(), "\n")[0]
		rowValues := strings.Split(strings.TrimSpace(firstRowValVector), "=>")
		value := strings.Split(strings.TrimSpace(rowValues[1]), " ")
		switch metric {
		case promCpuUsageAvg:
			cpuUsage, _ := strconv.ParseFloat(value[0], 64)
			nodeMetrics.Cpu = cpuUsage
		case promMemUsageAvg:
			memUsage, _ := strconv.ParseFloat(value[0], 64)
			nodeMetrics.Memory = memUsage
		}
	}
	return nodeMetrics, nil
}
