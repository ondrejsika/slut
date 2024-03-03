package random

import (
	"math/rand"
	"time"

	sleep_cmd "github.com/sikalabs/slu/cmd/sleep"
	"github.com/spf13/cobra"
)

var FlagMinTime int
var FlagMaxTime int
var FlagVerbose bool

var Cmd = &cobra.Command{
	Use:     "random",
	Short:   "Sleep random time",
	Aliases: []string{"r", "ran", "rnd"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		rand.Seed(time.Now().UnixNano())
		sleepTimeInt := rand.Intn(FlagMaxTime-FlagMinTime) + FlagMinTime
		if FlagVerbose {
			c.Println("Sleep for", sleepTimeInt, "ms")

		}
		time.Sleep(time.Duration(sleepTimeInt) * time.Millisecond)
	},
}

func init() {
	sleep_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().IntVar(
		&FlagMinTime,
		"min",
		0,
		"Minimum sleep time (in ms)",
	)
	Cmd.Flags().IntVar(
		&FlagMaxTime,
		"max",
		1000, // 1s
		"Maximum sleep time (in ms)",
	)
	Cmd.Flags().BoolVarP(
		&FlagVerbose,
		"verbose",
		"v",
		false,
		"verbose output",
	)
}
