package run

import "github.com/sunshinexcode/gotk/vcmd"

var (
	initTcpCmd = &vcmd.Command{
		Use:   "init-tcp",
		Short: "initialize new tcp project in current directory",
		Run:   initTcpProject,
	}
)

// initTcpProject init tcp project
func initTcpProject(_ *vcmd.Command, args []string) {
	initProject(args, CliTcp)
}
