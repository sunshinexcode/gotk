package vmetric

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/sunshinexcode/gotk/vlog"
	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vreflect"
)

type (
	Metric struct {
		BuildInfo           *prometheus.GaugeVec
		HttpRequestDuration *prometheus.HistogramVec
		HttpRequestTotal    *prometheus.CounterVec
		Options             *Options
	}

	Options struct {
		DisableGoCollector              bool              `mapstructure:",omitempty"`
		DisableProcessCollector         bool              `mapstructure:",omitempty"`
		BuildInfoConstLabels            map[string]string `mapstructure:",omitempty"`
		BuildInfoHelp                   string            `mapstructure:",omitempty"`
		BuildInfoLabelNames             []string          `mapstructure:",omitempty"`
		BuildInfoName                   string            `mapstructure:",omitempty"`
		HttpRequestDurationBucketsCount int               `mapstructure:",omitempty"`
		HttpRequestDurationBucketsStart float64           `mapstructure:",omitempty"`
		HttpRequestDurationBucketsWidth float64           `mapstructure:",omitempty"`
		HttpRequestDurationConstLabels  map[string]string `mapstructure:",omitempty"`
		HttpRequestDurationHelp         string            `mapstructure:",omitempty"`
		HttpRequestDurationLabelNames   []string          `mapstructure:",omitempty"`
		HttpRequestDurationName         string            `mapstructure:",omitempty"`
		HttpRequestTotalConstLabels     map[string]string `mapstructure:",omitempty"`
		HttpRequestTotalHelp            string            `mapstructure:",omitempty"`
		HttpRequestTotalLabelNames      []string          `mapstructure:",omitempty"`
		HttpRequestTotalName            string            `mapstructure:",omitempty"`
		Namespace                       string            `mapstructure:",omitempty"`
		Port                            string            `mapstructure:",omitempty"`
		Url                             string            `mapstructure:",omitempty"`
	}
)

const (
	MetricTypeApi           = "api"
	MetricTypeThirdPartyApi = "thirdPartyApi"
	MetricTypeCron          = "cron"
)

var (
	defaultOptions = map[string]any{
		"DisableGoCollector":              false,
		"DisableProcessCollector":         false,
		"BuildInfoConstLabels":            map[string]string{},
		"BuildInfoHelp":                   "build info",
		"BuildInfoLabelNames":             []string{},
		"BuildInfoName":                   "build_info",
		"HttpRequestDurationBucketsCount": 16,
		"HttpRequestDurationBucketsStart": float64(0),
		"HttpRequestDurationBucketsWidth": float64(200),
		"HttpRequestDurationConstLabels":  map[string]string{},
		"HttpRequestDurationHelp":         "http request duration",
		"HttpRequestDurationLabelNames":   []string{"api", "code"},
		"HttpRequestDurationName":         "http_request_duration",
		"HttpRequestTotalConstLabels":     map[string]string{},
		"HttpRequestTotalHelp":            "http request total",
		"HttpRequestTotalLabelNames":      []string{"api", "code"},
		"HttpRequestTotalName":            "http_request_total",
		"Namespace":                       "",
		"Port":                            ":9090",
		"Url":                             "/metrics",
	}
)

// Init init metric
func Init(options map[string]any) (metric *Metric, err error) {
	metric, err = New(options)
	if err != nil {
		return
	}

	// New metric
	metric.NewHttpRequestTotal()
	metric.NewHttpRequestDuration()
	metric.NewBuildInfo()

	// Disable
	if metric.Options.DisableGoCollector {
		prometheus.Unregister(collectors.NewGoCollector())
	}
	if metric.Options.DisableProcessCollector {
		prometheus.Unregister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	}

	return
}

// New create new metric
func New(options map[string]any) (metric *Metric, err error) {
	metric = &Metric{Options: &Options{}}
	err = metric.SetConfig(options)
	return
}

// D return HttpRequestDuration
func (metric *Metric) D() *prometheus.HistogramVec {
	return metric.HttpRequestDuration
}

// NewBuildInfo new build info
func (metric *Metric) NewBuildInfo() {
	metric.BuildInfo = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace:   metric.Options.Namespace,
		Name:        metric.Options.BuildInfoName,
		Help:        metric.Options.BuildInfoHelp,
		ConstLabels: metric.Options.BuildInfoConstLabels,
	}, metric.Options.BuildInfoLabelNames,
	)
}

// NewHttpRequestDuration new http request duration, millisecond
func (metric *Metric) NewHttpRequestDuration() {
	metric.HttpRequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: metric.Options.Namespace,
		Name:      metric.Options.HttpRequestDurationName,
		Help:      metric.Options.HttpRequestDurationHelp,
		Buckets: prometheus.LinearBuckets(metric.Options.HttpRequestDurationBucketsStart,
			metric.Options.HttpRequestDurationBucketsWidth, metric.Options.HttpRequestDurationBucketsCount),
		ConstLabels: metric.Options.HttpRequestDurationConstLabels,
	}, metric.Options.HttpRequestDurationLabelNames,
	)
}

// NewHttpRequestTotal new http request total
func (metric *Metric) NewHttpRequestTotal() {
	metric.HttpRequestTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace:   metric.Options.Namespace,
		Name:        metric.Options.HttpRequestTotalName,
		Help:        metric.Options.HttpRequestTotalHelp,
		ConstLabels: metric.Options.HttpRequestTotalConstLabels,
	}, metric.Options.HttpRequestTotalLabelNames,
	)
}

// Run http
func (metric *Metric) Run() (err error) {
	go func() {
		http.Handle(metric.Options.Url, promhttp.Handler())
		if err = http.ListenAndServe(metric.Options.Port, nil); err != nil {
			vlog.Fatal("start metric server", "err", err)
		}
	}()

	return
}

// SetConfig set config
func (metric *Metric) SetConfig(options map[string]any) error {
	return vreflect.SetAttrs(metric.Options, vmap.Merge(defaultOptions, options))
}

// T return HttpRequestTotal
func (metric *Metric) T() *prometheus.CounterVec {
	return metric.HttpRequestTotal
}

// Reset reset metric
func Reset() {
	prometheus.DefaultRegisterer = prometheus.NewRegistry()
}
