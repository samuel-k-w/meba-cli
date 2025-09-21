package cmd

import (
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update meba packages",
	Long:  `Update meba packages to the latest compatible versions.`,
	Run: func(cmd *cobra.Command, args []string) {
		updatePackages()
	},
}

func updatePackages() {
	color.Blue("🔄 Updating packages...")
	
	// Update go modules
	color.Blue("📦 Updating Go modules...")
	updateCmd := exec.Command("go", "get", "-u", "./...")
	updateCmd.Stdout = os.Stdout
	updateCmd.Stderr = os.Stderr

	if err := updateCmd.Run(); err != nil {
		color.Red("Failed to update packages: %v", err)
		os.Exit(1)
	}

	// Tidy modules
	color.Blue("🧹 Tidying modules...")
	tidyCmd := exec.Command("go", "mod", "tidy")
	tidyCmd.Stdout = os.Stdout
	tidyCmd.Stderr = os.Stderr

	if err := tidyCmd.Run(); err != nil {
		color.Red("Failed to tidy modules: %v", err)
		os.Exit(1)
	}

	color.Green("✅ Packages updated successfully!")
	
	// Show updated packages
	color.Blue("📋 Updated packages:")
	listCmd := exec.Command("go", "list", "-m", "-u", "all")
	listCmd.Stdout = os.Stdout
	listCmd.Stderr = os.Stderr
	listCmd.Run()
}

func init() {
	rootCmd.AddCommand(updateCmd)
}