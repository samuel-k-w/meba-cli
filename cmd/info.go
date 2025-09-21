package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show environment information",
	Long:  "Display meba version, Go version, installed packages, and system info",
	Run: func(cmd *cobra.Command, args []string) {
		showInfo()
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}

func showInfo() {
	fmt.Println("🔍 Meba Environment Information")
	fmt.Println("================================")
	
	// Meba version
	fmt.Printf("📦 Meba CLI Version: %s\n", getVersion())
	
	// Go version
	fmt.Printf("🐹 Go Version: %s\n", runtime.Version())
	
	// OS and Architecture
	fmt.Printf("💻 OS/Arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
	
	// Current working directory
	if cwd, err := os.Getwd(); err == nil {
		fmt.Printf("📁 Working Directory: %s\n", cwd)
	}
	
	// Go module info
	showGoModuleInfo()
	
	// Installed packages
	showInstalledPackages()
	
	// Development tools
	showDevTools()
}

func getVersion() string {
	// This would be set during build
	return "1.0.0" // TODO: Set from build flags
}

func showGoModuleInfo() {
	fmt.Println("\n📋 Go Module Information:")
	fmt.Println("-------------------------")
	
	// Check if go.mod exists
	if _, err := os.Stat("go.mod"); err == nil {
		// Get module name
		if content, err := os.ReadFile("go.mod"); err == nil {
			lines := strings.Split(string(content), "\n")
			if len(lines) > 0 && strings.HasPrefix(lines[0], "module ") {
				moduleName := strings.TrimSpace(strings.TrimPrefix(lines[0], "module "))
				fmt.Printf("📦 Module: %s\n", moduleName)
			}
		}
		
		// Get Go version from mod file
		cmd := exec.Command("go", "version", "-m", ".")
		if output, err := cmd.Output(); err == nil {
			fmt.Printf("🔧 Module Info:\n%s\n", string(output))
		}
	} else {
		fmt.Println("❌ No go.mod file found in current directory")
	}
}

func showInstalledPackages() {
	fmt.Println("\n📚 Installed Dependencies:")
	fmt.Println("--------------------------")
	
	// Get dependencies from go.mod
	cmd := exec.Command("go", "list", "-m", "all")
	if output, err := cmd.Output(); err == nil {
		lines := strings.Split(string(output), "\n")
		count := 0
		for _, line := range lines {
			if strings.TrimSpace(line) != "" && !strings.Contains(line, "=>") {
				if count < 20 { // Limit output
					fmt.Printf("  %s\n", line)
				}
				count++
			}
		}
		if count > 20 {
			fmt.Printf("  ... and %d more dependencies\n", count-20)
		}
		fmt.Printf("\n📊 Total Dependencies: %d\n", count-1) // -1 for main module
	} else {
		fmt.Println("❌ Could not retrieve dependency information")
	}
}

func showDevTools() {
	fmt.Println("\n🛠️  Development Tools:")
	fmt.Println("----------------------")
	
	tools := map[string]string{
		"wire":      "github.com/google/wire/cmd/wire",
		"air":       "github.com/cosmtrek/air",
		"gotestsum": "gotest.tools/gotestsum",
		"swag":      "github.com/swaggo/swag/cmd/swag",
		"mockgen":   "github.com/golang/mock/mockgen",
	}
	
	for tool, pkg := range tools {
		if _, err := exec.LookPath(tool); err == nil {
			// Get version if possible
			cmd := exec.Command(tool, "version")
			if output, err := cmd.Output(); err == nil {
				version := strings.TrimSpace(string(output))
				fmt.Printf("  ✅ %s: %s\n", tool, version)
			} else {
				fmt.Printf("  ✅ %s: installed\n", tool)
			}
		} else {
			fmt.Printf("  ❌ %s: not installed (go install %s@latest)\n", tool, pkg)
		}
	}
}