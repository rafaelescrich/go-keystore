package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Go Keystore",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Go Keytore - version 0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
