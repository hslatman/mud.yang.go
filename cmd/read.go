package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"

	"github.com/openconfig/ygot/ygot"

	"github.com/hslatman/mud.yang.go/pkg/mudyang"
)

var fileToRead string

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Reads a MUD file",
	Long:  `Reads and dumps contents of a MUD file`,
	Run: func(cmd *cobra.Command, args []string) {
		
		if fileToRead == "" {
			println("File to read not specified; use -f to specify the MUD to read")
			return
		}

	  	json, _ := ioutil.ReadFile(fileToRead)
		mud := &mudyang.Mudfile{}
		if err := mudyang.Unmarshal([]byte(json), mud); err != nil {
			println(fmt.Sprintf("Can't unmarshal JSON: %v", err))
			return
		}

		// TODO: print a summary instead of the full JSON?

		// println(*mud.Mud.MudUrl)
		// println(*mud.Mud.MudVersion)
		// println(mud.Mud.MudSignature)

		jsonString, err := ygot.EmitJSON(mud, &ygot.EmitJSONConfig{
			Format: ygot.RFC7951,
			Indent: "  ",
			RFC7951Config: &ygot.RFC7951JSONConfig{
				AppendModuleName: true,
			},
		})

		if err != nil {
			println(fmt.Sprintf("JSON error: %v", err))
			return
		}

		println(jsonString)
	},
}
