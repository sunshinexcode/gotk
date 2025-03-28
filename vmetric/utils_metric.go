package vmetric

import "time"

func MetricHttpRequestCount(metric *Metric, count float64, lvs ...string) {
	metric.HttpRequestTotal.WithLabelValues(lvs...).Add(count)
}

func MetricHttpRequestCountTypeApi(metric *Metric, count float64, lvs ...string) {
	MetricHttpRequestCount(metric, count, append([]string{MetricTypeApi}, lvs...)...)
}

func MetricHttpRequestCountTypeCron(metric *Metric, count float64, lvs ...string) {
	MetricHttpRequestCount(metric, count, append([]string{MetricTypeCron}, lvs...)...)
}

func MetricHttpRequestCountTypeThirdPartyApi(metric *Metric, count float64, lvs ...string) {
	MetricHttpRequestCount(metric, count, append([]string{MetricTypeThirdPartyApi}, lvs...)...)
}

// lvs: type/api/code
func MetricHttpRequestTotal(metric *Metric, lvs ...string) {
	metric.HttpRequestTotal.WithLabelValues(lvs...).Add(1)
}

func MetricHttpRequestTotalTypeApi(metric *Metric, lvs ...string) {
	MetricHttpRequestTotal(metric, append([]string{MetricTypeApi}, lvs...)...)
}

func MetricHttpRequestTotalTypeCron(metric *Metric, lvs ...string) {
	MetricHttpRequestTotal(metric, append([]string{MetricTypeCron}, lvs...)...)
}

func MetricHttpRequestTotalTypeThirdPartyApi(metric *Metric, lvs ...string) {
	MetricHttpRequestTotal(metric, append([]string{MetricTypeThirdPartyApi}, lvs...)...)
}

// lvs: type/api/code
func MetricHttpRequestDuration(metric *Metric, timeStart time.Time, lvs ...string) {
	metric.HttpRequestDuration.WithLabelValues(lvs...).Observe(float64(time.Since(timeStart).Milliseconds()))
}

func MetricHttpRequestDurationTypeApi(metric *Metric, timeStart time.Time, lvs ...string) {
	MetricHttpRequestDuration(metric, timeStart, append([]string{MetricTypeApi}, lvs...)...)
}

func MetricHttpRequestDurationTypeCron(metric *Metric, timeStart time.Time, lvs ...string) {
	MetricHttpRequestDuration(metric, timeStart, append([]string{MetricTypeCron}, lvs...)...)
}

func MetricHttpRequestDurationTypeThirdPartyApi(metric *Metric, timeStart time.Time, lvs ...string) {
	MetricHttpRequestDuration(metric, timeStart, append([]string{MetricTypeThirdPartyApi}, lvs...)...)
}
