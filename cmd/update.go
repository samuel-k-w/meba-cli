package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update packages to latest versions",
	Long:  "Upgrade meba packages to the latest compatible versions",
	Run: func(cmd *cobra.Command, args []string) {
		updatePackages()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func updatePackages() {
	fmt.Println("🔄 Updating packages to latest versions...")
	
	// Check if go.mod exists
	if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
		fmt.Println("❌ No go.mod file found. Run this command in a Go module directory.")
		os.Exit(1)
	}
	
	// Get current dependencies
	fmt.Println("📋 Checking current dependencies...")
	
	// Update all dependencies
	fmt.Println("⬆️  Updating all dependencies...")
	cmd := exec.Command("go", "get", "-u", "./...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	if err := cmd.Run(); err != nil {
		fmt.Printf("❌ Failed to update dependencies: %v\n", err)
		os.Exit(1)
	}
	
	// Run go mod tidy
	fmt.Println("🧹 Cleaning up dependencies...")
	cmd = exec.Command("go", "mod", "tidy")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	if err := cmd.Run(); err != nil {
		fmt.Printf("❌ Failed to tidy dependencies: %v\n", err)
		os.Exit(1)
	}
	
	// Update development tools
	updateDevTools()
	
	// Show updated packages
	showUpdatedPackages()
	
	fmt.Println("✅ All packages updated successfully!")
	fmt.Println("💡 Run 'meba test' to ensure everything still works")
}

func updateDevTools() {
	fmt.Println("🛠️  Updating development tools...")
	
	tools := []string{
		"github.com/google/wire/cmd/wire@latest",
		"github.com/cosmtrek/air@latest",
		"gotest.tools/gotestsum@latest",
		"github.com/swaggo/swag/cmd/swag@latest",
		"github.com/golang/mock/mockgen@latest",
	}
	
	for _, tool := range tools {
		fmt.Printf("  📦 Updating %s...\n", strings.Split(tool, "@")[0])
		cmd := exec.Command("go", "install", tool)
		if err := cmd.Run(); err != nil {
			fmt.Printf("    ⚠️  Warning: Could not update %s: %v\n", tool, err)
		} else {
			fmt.Printf("    ✅ Updated %s\n", strings.Split(tool, "@")[0])
		}
	}
}

func showUpdatedPackages() {
	fmt.Println("\n📊 Updated Dependencies:")
	fmt.Println("------------------------")
	
	cmd := exec.Command("go", "list", "-m", "-u", "all")
	if output, err := cmd.Output(); err == nil {
		lines := strings.Split(string(output), "\n")
		updatedCount := 0
		
		for _, line := range lines {
			if strings.Contains(line, "[") && strings.Contains(line, "]") {
				fmt.Printf("  📈 %s\n", line)
				updatedCount++
			}
		}
		
		if updatedCount == 0 {
			fmt.Println("  ✅ All dependencies were already up to date")
		} else {
			fmt.Printf("\n📊 Total packages updated: %d\n", updatedCount)
		}
	} else {
		fmt.Println("❌ Could not retrieve update information")
	}
}