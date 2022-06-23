package appdmetricsexporter

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

type metricsexporter struct {
	// Input configuration.
	config *Config
	logger       *zap.Logger
	appdClient   *http.Client
}

func (e *metricsexporter) start(ctx context.Context, host component.Host) (err error) {
	var timeout = time.Second * 10
	if (e.config.TimeoutSettings.Timeout.String() != ""){
		timeout = time.Duration(e.config.TimeoutSettings.Timeout.Seconds())
	}
	// e.config.TimeoutSettings.Timeout.Milliseconds()
	e.appdClient = &http.Client{
		Timeout: timeout,
	}

	return
}

func (e *metricsexporter) shutdown(ctx context.Context) (err error) {
	e.appdClient.CloseIdleConnections()
	return
}

func (e *metricsexporter) pushMetrics(ctx context.Context, metrics pmetric.Metrics) error {
	endpoint := "http://" + e.config.MachineAgentHost + ":" + strconv.Itoa(e.config.MachineAgentPort) + "/api/v1/metrics"
	metricsJson, _ := json.Marshal(generateAppDMetrics(metrics))
	resp, error := e.appdClient.Post(endpoint, "application/json", bytes.NewBuffer(metricsJson))
	defer resp.Body.Close()

	return error
}