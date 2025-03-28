package bootstrap

import (
	"github.com/sunshinexcode/gotk/vbootstrap"

	"app/configs"
)

// NewLog new log
func NewLog(config *configs.Config) (err error) {
	return vbootstrap.NewLog(config)
}
