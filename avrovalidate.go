package main

import (
	"github.com/hekonsek/osexit"
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "avro-validate",
	Run: func(cmd *cobra.Command, args []string) {
		osexit.ExitOnError(cmd.Help())
	},
}

func main() {
	osexit.ExitOnError(rootCommand.Execute())
}
