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
		println(err)
		os.Exit(1)
	}
}

func init() {

	readCmd.PersistentFlags().StringVarP(&fileToRead, "file", "f", "", "MUD file to read")
	readCmd.MarkFlagRequired("file")

	rootCmd.AddCommand(readCmd)
	rootCmd.AddCommand(validateCmd)

}
