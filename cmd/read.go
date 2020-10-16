package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"

	"github.com/openconfig/ygot/ygot"

	"github.com/hslatman/mud.yang.go/pkg/mudyang"
)

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Reads a MUD file",
	Long:  `Reads and dumps contents of a MUD file`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		fileToRead := args[0]

		json, err := ioutil.ReadFile(fileToRead)
		if err != nil {
			println(fmt.Sprintf("File could not be read: %v", err))
			return
		}

		mud := &mudyang.Mudfile{}
		if err := mudyang.Unmarshal([]byte(json), mud); err != nil {
			println(fmt.Sprintf("Can't unmarshal JSON: %v", err))
			return
		}

		// TODO: print a summary instead of the full JSON?

		jsonString, err := ygot.EmitJSON(mud, &ygot.EmitJSONConfig{
			Format: ygot.RFC7951,
			Indent: "  ",
			RFC7951Config: &ygot.RFC7951JSONConfig{
				AppendModuleName: true,
			},
			SkipValidation: false,
		})

		if err != nil {
			println(fmt.Sprintf("JSON error: %v", err))
			return
		}

		println(jsonString)
	},
}
