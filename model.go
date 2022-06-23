package appdmetricsexporter

import (
	"strings"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

type AppdCustomMetric struct {
	MetricName string `json:"metricName"`
    AggregationType string `json:"aggregatorType"`
    Value float64 `json:"value"`
}

func generateAppDMetrics(metric pmetric.Metrics) []AppdCustomMetric{
    resourceMetrics := metric.ResourceMetrics()
    appdMetricsSlice := []AppdCustomMetric{}

    for irm := 1; irm < resourceMetrics.Len(); irm++{
		rm := resourceMetrics.At(irm)
		scopeMetrics := rm.ScopeMetrics()
		for ism := 1; ism < scopeMetrics.Len(); ism++{
			sm := scopeMetrics.At(ism)
			metrics := sm.Metrics()
			for im := 1; im < scopeMetrics.Len(); im++{
				m := metrics.At(im)
				appdMetric := &AppdCustomMetric{}
				if (m.DataType() == pdata.MetricDataTypeSum || m.DataType() == pdata.MetricDataTypeGauge){
					appdMetric.AggregationType = "SUM"
					switch m.DataType() {
						case  pdata.MetricDataTypeSum:
							metricValue := m.Sum().DataPoints().At(0)
							appdMetric.Value = metricValue.DoubleVal()
						case pdata.MetricDataTypeGauge:
							metricValue := m.Gauge().DataPoints().At(0)
							appdMetric.Value = metricValue.DoubleVal()
					}				                
					appdMetricPath := "OpenTelemetry | " + strings.ReplaceAll(m.Name(), ".", " | ")					
					appdMetric.MetricName = appdMetricPath
					appdMetricsSlice = append(appdMetricsSlice, *appdMetric)
				}
			}

		}
	}
	return appdMetricsSlice
}