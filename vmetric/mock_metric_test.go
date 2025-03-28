package vmetric_test

import (
	"testing"

	dto "github.com/prometheus/client_model/go"

	"github.com/sunshinexcode/gotk/vmetric"
	"github.com/sunshinexcode/gotk/vmock"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestMock(t *testing.T) {
	metric, patches := vmetric.Mock()
	defer vmock.ResetAll(patches)

	metric.BuildInfo.WithLabelValues().Set(1)
	metric.HttpRequestTotal.WithLabelValues().Add(1)
	metric.HttpRequestDuration.WithLabelValues().Observe(1)

	vtest.Nil(t, metric.BuildInfo.WithLabelValues().Write(&dto.Metric{}))
	vtest.Nil(t, metric.HttpRequestTotal.WithLabelValues().Write(&dto.Metric{}))
	vtest.Equal(t, "Desc{fqName: \"\", help: \"\", constLabels: {}, variableLabels: {}}", metric.BuildInfo.WithLabelValues().Desc().String())
	vtest.Equal(t, "Desc{fqName: \"\", help: \"\", constLabels: {}, variableLabels: {}}", metric.HttpRequestTotal.WithLabelValues().Desc().String())
}
