package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/hslatman/mud.yang.go/pkg/mudyang"
	"github.com/openconfig/ygot/ytypes"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validates a MUD file",
	Long:  `Validates a MUD file`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		fileToRead := args[0]

		json, err := ioutil.ReadFile(fileToRead)
		if err != nil {
			println(err)
			return
		}

		mud := &mudyang.Mudfile{}
		if err := mudyang.Unmarshal([]byte(json), mud); err != nil {
			println(fmt.Sprintf("Can't unmarshal JSON: %v", err))
			return
		}

		options := &ytypes.LeafrefOptions{
			IgnoreMissingData: false,
			Log:               true,
		}
		if err = mud.Validate(options); err != nil {
			println(err)
			return
		}

		println("MUD file is valid")
	},
}
