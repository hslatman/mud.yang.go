package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mudyang",
	Short: "todo",
	Long:  `todo long`,
}

/*
Execute invokes the root Cobra command
*/
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(readCmd)
	rootCmd.AddCommand(validateCmd)
}
