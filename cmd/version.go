package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Meba CLI v%s\n", getVersion())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}