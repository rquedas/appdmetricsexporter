package appdmetricsexporter

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

const (
	typeStr = "appdmetrics"
	defaultPort = 8293
)

func createDefaultConfig() config.Exporter {
	return &Config{
		ExporterSettings:   config.NewExporterSettings(config.NewComponentID(typeStr)),
		TimeoutSettings:  exporterhelper.NewDefaultTimeoutSettings(),
		MachineAgentPort: defaultPort,
	}
}

func createMetricsExporter(
	_ context.Context,
	set component.ExporterCreateSettings,
	cfg config.Exporter,
) (component.MetricsExporter, error) {
	appdExporter, err := newExporter(cfg, set)
	if err != nil {
		return nil, err
	}
	oCfg := cfg.(*Config)
	return exporterhelper.NewMetricsExporter(
		cfg,
		set,
		appdExporter.pushMetrics,
		exporterhelper.WithCapabilities(consumer.Capabilities{MutatesData: false}),
		exporterhelper.WithTimeout(oCfg.TimeoutSettings),
		exporterhelper.WithStart(appdExporter.start),
		exporterhelper.WithShutdown(appdExporter.shutdown),
	)
}

func newExporter(cfg config.Exporter, set component.ExporterCreateSettings) (*metricsexporter, error){
	logger := set.Logger
	metricsExporterCfg := cfg.(*Config)

	appdExporter := &metricsexporter{
		logger:       logger,
		config:       metricsExporterCfg,
	}
	
	return appdExporter, nil

}

// NewFactory creates a factory for OTLP exporter.
func NewFactory() component.ExporterFactory {
	return component.NewExporterFactory(
		typeStr,
		createDefaultConfig,
		component.WithMetricsExporter(createMetricsExporter),
	)
}