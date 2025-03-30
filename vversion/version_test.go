package vversion_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vmetric"
	"github.com/sunshinexcode/gotk/vmock"
	"github.com/sunshinexcode/gotk/vtest"
	"github.com/sunshinexcode/gotk/vversion"
)

func TestGotkVersion(t *testing.T) {
	vtest.Equal(t, "v0.0.2", vversion.GotkVersion)
}

func TestMetric(t *testing.T) {
	metric, patches := vmetric.Mock()
	defer vmock.ResetAll(patches)

	vversion.Metric(metric)
}

func TestPrint(t *testing.T) {
	vversion.Print()
}
