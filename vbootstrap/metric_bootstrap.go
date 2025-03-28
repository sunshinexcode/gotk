package vbootstrap

import (
	"github.com/sunshinexcode/gotk/vconfig"
	"github.com/sunshinexcode/gotk/vlog"
	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vmetric"
	"github.com/sunshinexcode/gotk/vversion"
)

// NewMetric new metric
func NewMetric(config vconfig.IConfig) (metric *vmetric.Metric, err error) {
	var options vmap.M
	metric = &vmetric.Metric{}

	defer func() {
		if err != nil {
			vlog.Error("NewMetric", "err", err, "options", metric.Options)
		}
	}()

	if err = vmap.Decode(config.GetMetric(), &options); err != nil {
		return
	}
	if metric, err = vmetric.Init(options); err != nil {
		return
	}
	if err = metric.Run(); err != nil {
		return
	}

	vversion.Metric(metric)
	vversion.Print()

	vlog.Infof("NewMetric, options:%+v", metric.Options)
	return
}
