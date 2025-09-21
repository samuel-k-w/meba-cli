package cmd

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display environment information",
	Long:  `Display information about the meba CLI, Go version, and installed packages.`,
	Run: func(cmd *cobra.Command, args []string) {
		displayInfo()
	},
}

func displayInfo() {
	color.Cyan("🔍 Meba CLI Information")
	fmt.Println(strings.Repeat("=", 40))
	
	// Meba version
	color.Blue("Meba CLI Version: ") 
	fmt.Println("1.0.0")
	
	// Go version
	color.Blue("Go Version: ")
	if output, err := exec.Command("go", "version").Output(); err == nil {
		fmt.Println(strings.TrimSpace(string(output)))
	} else {
		fmt.Println("Go not found")
	}
	
	// OS Info
	color.Blue("Operating System: ")
	fmt.Printf("%s %s\n", runtime.GOOS, runtime.GOARCH)
	
	// Node version (if available)
	color.Blue("Node Version: ")
	if output, err := exec.Command("node", "--version").Output(); err == nil {
		fmt.Println(strings.TrimSpace(string(output)))
	} else {
		fmt.Println("Node.js not found")
	}
	
	// Package Manager
	color.Blue("Package Manager: ")
	fmt.Println("Go Modules")
	
	fmt.Println()
	color.Cyan("📦 Installed Packages")
	fmt.Println(strings.Repeat("=", 40))
	
	// List go modules
	if output, err := exec.Command("go", "list", "-m", "all").Output(); err == nil {
		lines := strings.Split(string(output), "\n")
		for i, line := range lines {
			if i > 10 { // Limit output
				fmt.Printf("... and %d more packages\n", len(lines)-i)
				break
			}
			if strings.TrimSpace(line) != "" {
				fmt.Println(line)
			}
		}
	} else {
		fmt.Println("No go.mod found in current directory")
	}
}

func init() {
	rootCmd.AddCommand(infoCmd)
}