package ip

import (
	"os"
	"os/exec"
	"strconv"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/sikalabs/slu/utils/time_utils"
	"github.com/spf13/cobra"
)

var FlagJson bool
var FlagImage string
var FlagShell string

var Cmd = &cobra.Command{
	Use:   "kdev",
	Short: "Run sikalabs/dev in Kubernetes",
	Args:  cobra.MaximumNArgs(1),
	Run: func(c *cobra.Command, args []string) {
		cmd := exec.Command(
			"kubectl", "run",
			"dev-"+strconv.Itoa(time_utils.Unix()),
			"--rm", "-ti",
			"--image", FlagImage,
			"--", FlagShell,
		)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Run()
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagImage,
		"image",
		"i",
		"sikalabs/dev",
		"Container Image",
	)
	Cmd.Flags().StringVarP(
		&FlagShell,
		"shell",
		"s",
		"zsh",
		"Shell to run in container",
	)
}
