package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/openconfig/ygot/ygot"

	"github.com/hslatman/mud.yang.go/pkg/mudyang"
)

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Reads a MUD file",
	Long:  `Reads and dumps contents of a MUD file`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		fileToRead := args[0]

		json, err := os.ReadFile(fileToRead)
		if err != nil {
			return fmt.Errorf("failed to read file: %w", err)
		}

		mud := &mudyang.Mudfile{}
		if err := mudyang.Unmarshal([]byte(json), mud); err != nil {
			return fmt.Errorf("failed unmarshaling JSON: %w", err)
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
			return fmt.Errorf("failed to emit JSON: %w", err)
		}

		println(jsonString)

		return nil
	},
}
