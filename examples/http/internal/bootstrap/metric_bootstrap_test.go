package bootstrap_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vconfig"
	"github.com/sunshinexcode/gotk/vtest"

	"app/configs"
	"app/internal/bootstrap"
)

func TestNewMetric(t *testing.T) {
	vconfig.SetConfigPath("../../configs")

	config, err := configs.New()

	vtest.Nil(t, err)

	_, err = bootstrap.NewMetric(config)

	vtest.Nil(t, err)
}
