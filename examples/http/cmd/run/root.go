package run

import (
	"os"

	"github.com/sunshinexcode/gotk/vcmd"
	"github.com/sunshinexcode/gotk/vlog"

	"app/configs"
)

var (
	Version = "1.0.0"

	rootCmd = &vcmd.Command{
		Use:     "gotk",
		Version: Version,
		Short:   "Gotk Http Management CLI",
	}
)

func init() {
	vcmd.OnInitialize(configs.InitConfig)
	rootCmd.AddCommand(abCmd, cronCmd, httpCmd, versionCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		vlog.Error("execute", "err", err)
		os.Exit(1)
	}
}
