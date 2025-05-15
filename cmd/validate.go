package cmd

import (
	"fmt"
	"os"

	"github.com/openconfig/ygot/ytypes"
	"github.com/spf13/cobra"

	"github.com/hslatman/go-mudyang"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validates a MUD file",
	Long:  `Validates a MUD file`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fileToRead := args[0]

		json, err := os.ReadFile(fileToRead)
		if err != nil {
			return fmt.Errorf("failed to read file: %w", err)
		}

		mud := &mudyang.Mudfile{}
		if err := mudyang.Unmarshal([]byte(json), mud); err != nil {
			return fmt.Errorf("failed to unmarshal JSON: %w", err)
		}

		options := &ytypes.LeafrefOptions{
			IgnoreMissingData: false,
			Log:               true,
		}
		if err = mud.ΛValidate(options); err != nil {
			return fmt.Errorf("MUD validation failed: %w", err)
		}

		println("MUD file is valid")

		return nil
	},
}
