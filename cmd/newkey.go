package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var newKeyCmd = &cobra.Command{
	Use:   "new",
	Short: "Generate new key",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Ney key: " + strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(newKeyCmd)
}
