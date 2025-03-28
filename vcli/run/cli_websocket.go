package run

import "github.com/sunshinexcode/gotk/vcmd"

var (
	initWebsocketCmd = &vcmd.Command{
		Use:   "init-websocket",
		Short: "initialize new websocket project in current directory",
		Run:   initWebsocketProject,
	}
)

// initWebsocketProject init websocket project
func initWebsocketProject(_ *vcmd.Command, args []string) {
	initProject(args, CliWebsocket)
}
