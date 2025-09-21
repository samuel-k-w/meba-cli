package cmd

import (
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var buildWatch bool

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the meba application",
	Long:  `Build the meba application for production deployment.`,
	Run: func(cmd *cobra.Command, args []string) {
		if buildWatch {
			buildWithWatch()
		} else {
			buildProduction()
		}
	},
}

func buildProduction() {
	color.Blue("🔨 Building application...")
	
	// Create bin directory
	os.MkdirAll("bin", 0755)
	
	cmd := exec.Command("go", "build", "-o", "bin/server", "./cmd/server")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		color.Red("Build failed: %v", err)
		os.Exit(1)
	}

	color.Green("✅ Build completed successfully!")
	color.Blue("📦 Binary created at: bin/server")
}

func buildWithWatch() {
	color.Blue("🔨 Building with watch mode...")
	// This would require a file watcher implementation
	// For now, just do a regular build
	buildProduction()
}

func init() {
	rootCmd.AddCommand(buildCmd)
	buildCmd.Flags().BoolVar(&buildWatch, "watch", false, "Continuous compilation")
}