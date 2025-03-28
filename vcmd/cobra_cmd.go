package vcmd

import "github.com/spf13/cobra"

type (
	Command = cobra.Command
)

func OnInitialize(y ...func()) {
	cobra.OnInitialize(y...)
}
