package run

import "github.com/sunshinexcode/gotk/vcmd"

var (
	initHttpCmd = &vcmd.Command{
		Use:   "init",
		Short: "initialize new http project in current directory",
		Run:   initHttpProject,
	}
)

// initHttpProject init http project
func initHttpProject(_ *vcmd.Command, args []string) {
	initProject(args, CliHttp)
}
