package bootstrap

import (
	"github.com/sunshinexcode/gotk/vbootstrap"
	"github.com/sunshinexcode/gotk/ves"

	"app/configs"
)

// NewEs new es
func NewEs(config *configs.Config) (es *ves.Es, err error) {
	return vbootstrap.NewEs(config)
}
