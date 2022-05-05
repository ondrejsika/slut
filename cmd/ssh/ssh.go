package ssh

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "ssh",
	Short: "SSH Utils",
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
