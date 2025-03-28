package bootstrap

import (
	"github.com/sunshinexcode/gotk/vbootstrap"
	"github.com/sunshinexcode/gotk/vmetric"

	"app/configs"
)

// NewMetric new metric
func NewMetric(config *configs.Config) (metric *vmetric.Metric, err error) {
	return vbootstrap.NewMetric(config)
}
