package cmd

import "github.com/spf13/cobra"

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validates a MUD file",
	Long:  `Validates a MUD file`,
	Run: func(cmd *cobra.Command, args []string) {
		println("to be implemented")
	},
}
