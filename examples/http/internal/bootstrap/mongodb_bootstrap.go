package bootstrap

import (
	"github.com/sunshinexcode/gotk/vbootstrap"
	"github.com/sunshinexcode/gotk/vmongodb"

	"app/configs"
)

// NewMongodb new mongodb
func NewMongodb(config *configs.Config) (mongodb *vmongodb.Mongodb, err error) {
	return vbootstrap.NewMongodb(config)
}
