package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/hekonsek/osexit"
	"github.com/linkedin/goavro"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var jsonCommand = &cobra.Command{
	Use: "json schemafile jsonfile",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			osexit.ExitOnError(cmd.Help())
			os.Exit(osexit.UnixExitCodeGeneralError)
		}

		// Read schema file
		schema, err := ioutil.ReadFile(args[0])
		osexit.ExitOnError(err)
		codec, err := goavro.NewCodec(string(schema))
		osexit.ExitOnError(err)

		// Read JSON file
		jsonData, err := ioutil.ReadFile(args[1])
		osexit.ExitOnError(err)
		var jsonObject map[string]interface{}
		err = json.Unmarshal(jsonData, &jsonObject)
		osexit.ExitOnError(err)

		_, err = codec.BinaryFromNative(nil, jsonObject)
		if err != nil {
			fmt.Printf("%s. JSON file %s doesn't comply to schema file %s: . %s\n",
				color.RedString("Failure"), color.YellowString(args[0]), color.YellowString(args[1]), err.Error())
		} else {
			fmt.Printf("%s. JSON file %s complies to schema file %s.\n",
				color.GreenString("Success"), color.YellowString(args[0]), color.YellowString(args[1]))
		}
	},
}

func init() {
	rootCommand.AddCommand(jsonCommand)
}
