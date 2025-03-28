package run

import (
	"os"

	"github.com/sunshinexcode/gotk/vcmd"
	"github.com/sunshinexcode/gotk/vlog"
)

var (
	Version = "0.0.1"

	rootCmd = &vcmd.Command{
		Use:     "gotk",
		Version: Version,
		Short:   "Gotk Management CLI",
	}
)

func init() {
	rootCmd.AddCommand(initHttpCmd, initTcpCmd, initWebsocketCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		vlog.Error("execute", "err", err)
		os.Exit(1)
	}
}
