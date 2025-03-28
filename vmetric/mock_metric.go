package vmetric

import (
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"

	"github.com/sunshinexcode/gotk/vmock"
)

type MockGauge struct {
}

func (m *MockGauge) Add(float64) {
}

func (m *MockGauge) Collect(chan<- prometheus.Metric) {
}

func (m *MockGauge) Dec() {
}

func (m *MockGauge) Desc() *prometheus.Desc {
	return &prometheus.Desc{}
}

func (m *MockGauge) Describe(chan<- *prometheus.Desc) {
}

func (m *MockGauge) Inc() {
}

func (m *MockGauge) Set(float64) {
}

func (m *MockGauge) SetToCurrentTime() {
}

func (m *MockGauge) Sub(float64) {
}

func (m *MockGauge) Write(*dto.Metric) error {
	return nil
}

type MockCounter struct {
}

func (m *MockCounter) Add(float64) {
}

func (m *MockCounter) Collect(chan<- prometheus.Metric) {
}

func (m *MockCounter) Desc() *prometheus.Desc {
	return &prometheus.Desc{}
}

func (m *MockCounter) Describe(chan<- *prometheus.Desc) {
}

func (m *MockCounter) Inc() {
}

func (m *MockCounter) Write(*dto.Metric) error {
	return nil
}

type MockObserver struct {
}

func (m *MockObserver) Observe(float64) {
}

func Mock() (metric *Metric, patches []*vmock.Patches) {
	metric = &Metric{BuildInfo: &prometheus.GaugeVec{}, HttpRequestTotal: &prometheus.CounterVec{}, HttpRequestDuration: &prometheus.HistogramVec{}}

	patchGaugeVec := vmock.ApplyMethodReturn(&prometheus.GaugeVec{}, "WithLabelValues", &MockGauge{})
	patches = append(patches, patchGaugeVec)

	patchCounterVec := vmock.ApplyMethodReturn(&prometheus.CounterVec{}, "WithLabelValues", &MockCounter{})
	patches = append(patches, patchCounterVec)

	patchHistogramVec := vmock.ApplyMethodReturn(&prometheus.HistogramVec{}, "WithLabelValues", &MockObserver{})
	patches = append(patches, patchHistogramVec)

	return
}
