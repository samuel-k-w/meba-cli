package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

var (
	buildWatchFlag bool
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the application",
	Long:  "Compile the Go project to executable",
	Run: func(cmd *cobra.Command, args []string) {
		if buildWatchFlag {
			buildWithWatch()
		} else {
			if err := buildApp(); err != nil {
				fmt.Printf("❌ Build failed: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("✅ Build completed successfully!")
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	buildCmd.Flags().BoolVarP(&buildWatchFlag, "watch", "w", false, "Continuous compilation")
}

func buildApp() error {
	fmt.Println("🔨 Building application...")
	
	// Create dist directory
	if err := os.MkdirAll("dist", 0755); err != nil {
		return fmt.Errorf("failed to create dist directory: %w", err)
	}
	
	// Auto-generate wire before building
	fmt.Println("🔧 Generating wire dependencies...")
	wireCmd := exec.Command("wire", "./internal")
	wireCmd.Run() // Don't fail if wire fails
	
	// Build the application
	outputPath := filepath.Join("dist", "server")
	cmd := exec.Command("go", "build", "-o", outputPath, "./cmd/server")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	return cmd.Run()
}

func buildWithWatch() {
	fmt.Println("🔄 Starting continuous compilation...")
	
	// Initial build
	if err := buildApp(); err != nil {
		fmt.Printf("❌ Initial build failed: %v\n", err)
	}
	
	// Watch for changes (simple implementation)
	fmt.Println("👀 Watching for changes... (Press Ctrl+C to stop)")
	
	lastBuild := time.Now()
	for {
		time.Sleep(2 * time.Second)
		
		// Check if any .go files have been modified
		if shouldRebuild(lastBuild) {
			fmt.Println("📝 Changes detected, rebuilding...")
			if err := buildApp(); err != nil {
				fmt.Printf("❌ Build failed: %v\n", err)
			} else {
				fmt.Println("✅ Build completed!")
			}
			lastBuild = time.Now()
		}
	}
}

func shouldRebuild(lastBuild time.Time) bool {
	// Simple file modification check
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		
		// Skip vendor, node_modules, dist, tmp directories
		if info.IsDir() {
			name := info.Name()
			if name == "vendor" || name == "node_modules" || name == "dist" || name == "tmp" || name == ".git" {
				return filepath.SkipDir
			}
			return nil
		}
		
		// Check .go files
		if filepath.Ext(path) == ".go" && info.ModTime().After(lastBuild) {
			return fmt.Errorf("rebuild needed")
		}
		
		return nil
	})
	
	return err != nil
}