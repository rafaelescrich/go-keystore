package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-keystore",
	Short: "Go Keystore is a storage for your keys",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: verify if first time (hasDB?)
		// then if not first time asks for password
	},
}

// Execute is the function that starts the cli
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
