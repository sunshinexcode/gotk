package vmetric_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vmetric"
	"github.com/sunshinexcode/gotk/vmock"
	"github.com/sunshinexcode/gotk/vtest"
	"github.com/sunshinexcode/gotk/vtime"
)

func TestMetricHttpRequestTotalMock(t *testing.T) {
	_, patches := vmetric.Mock()
	defer vmock.ResetAll(patches)

	metric, err := vmetric.New(vmap.M{"HttpRequestTotalLabelNames": []string{"type", "api", "code"}})

	vtest.Nil(t, err)

	vmetric.MetricHttpRequestTotal(metric, vmetric.MetricTypeApi, "api", "code")
}

func TestMetricHttpRequestTotalTypeApiMock(t *testing.T) {
	_, patches := vmetric.Mock()
	defer vmock.ResetAll(patches)

	metric, err := vmetric.New(vmap.M{"HttpRequestTotalLabelNames": []string{"type", "api", "code"}})

	vtest.Nil(t, err)

	vmetric.MetricHttpRequestTotalTypeApi(metric, "api", "code")
}

func TestMetricHttpRequestTotalTypeCronMock(t *testing.T) {
	_, patches := vmetric.Mock()
	defer vmock.ResetAll(patches)

	metric, err := vmetric.New(vmap.M{"HttpRequestTotalLabelNames": []string{"type", "api", "code"}})

	vtest.Nil(t, err)

	vmetric.MetricHttpRequestTotalTypeCron(metric, "api", "code")
}

func TestMetricHttpRequestTotalTypeThirdPartyApiMock(t *testing.T) {
	_, patches := vmetric.Mock()
	defer vmock.ResetAll(patches)

	metric, err := vmetric.New(vmap.M{"HttpRequestTotalLabelNames": []string{"type", "api", "code"}})

	vtest.Nil(t, err)

	vmetric.MetricHttpRequestTotalTypeThirdPartyApi(metric, "api", "code")
}

func TestMetricHttpRequestDurationMock(t *testing.T) {
	_, patches := vmetric.Mock()
	defer vmock.ResetAll(patches)

	metric, err := vmetric.New(vmap.M{"HttpRequestDurationLabelNames": []string{"type", "api", "code"}})

	vtest.Nil(t, err)

	vmetric.MetricHttpRequestDuration(metric, vtime.GetNow(), vmetric.MetricTypeApi, "api", "code")
}

func TestMetricHttpRequestDurationTypeApiMock(t *testing.T) {
	_, patches := vmetric.Mock()
	defer vmock.ResetAll(patches)

	metric, err := vmetric.New(vmap.M{"HttpRequestDurationLabelNames": []string{"type", "api", "code"}})

	vtest.Nil(t, err)

	vmetric.MetricHttpRequestDurationTypeApi(metric, vtime.GetNow(), "api", "code")
}

func TestMetricHttpRequestDurationTypeCronMock(t *testing.T) {
	_, patches := vmetric.Mock()
	defer vmock.ResetAll(patches)

	metric, err := vmetric.New(vmap.M{"HttpRequestDurationLabelNames": []string{"type", "api", "code"}})

	vtest.Nil(t, err)

	vmetric.MetricHttpRequestDurationTypeCron(metric, vtime.GetNow(), "api", "code")
}

func TestMetricHttpRequestDurationTypeThirdPartyApiMock(t *testing.T) {
	_, patches := vmetric.Mock()
	defer vmock.ResetAll(patches)

	metric, err := vmetric.New(vmap.M{"HttpRequestDurationLabelNames": []string{"type", "api", "code"}})

	vtest.Nil(t, err)

	vmetric.MetricHttpRequestDurationTypeThirdPartyApi(metric, vtime.GetNow(), "api", "code")
}
