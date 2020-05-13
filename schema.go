package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/hekonsek/osexit"
	"github.com/linkedin/goavro"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var schemaCommand = &cobra.Command{
	Use:   "schema schemafile [schemafile...]",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			osexit.ExitOnError(cmd.Help())
			os.Exit(osexit.UnixExitCodeGeneralError)
		}

		for _, schemaFile := range args {
			schema, err := ioutil.ReadFile(schemaFile)
			osexit.ExitOnError(err)

			var status string
			_, err = goavro.NewCodec(string(schema))
			if err != nil {
				status = fmt.Sprintf("%s (%s)", color.RedString("INVALID"), err.Error())
			} else {
				status = color.GreenString("OK")
			}
			fmt.Printf("%s: %s\n", color.YellowString(schemaFile), status)
		}
	},
}

func init() {
	rootCommand.AddCommand(schemaCommand)
}

