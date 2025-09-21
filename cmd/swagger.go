package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var swaggerCmd = &cobra.Command{
	Use:   "swagger",
	Short: "Generate Swagger documentation",
	Long:  `Generate Swagger/OpenAPI documentation for your API endpoints`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := generateSwagger(); err != nil {
			color.Red("Error generating Swagger docs: %v", err)
			os.Exit(1)
		}
		color.Green("✅ Swagger documentation generated successfully!")
		fmt.Println("📚 Docs available at: http://localhost:8080/swagger/index.html")
	},
}

func generateSwagger() error {
	// Check if swag is installed
	if _, err := exec.LookPath("swag"); err != nil {
		color.Yellow("⚠️  swag CLI not found, installing...")
		installCmd := exec.Command("go", "install", "github.com/swaggo/swag/cmd/swag@latest")
		if err := installCmd.Run(); err != nil {
			return fmt.Errorf("failed to install swag CLI: %w", err)
		}
		color.Green("✅ swag CLI installed")
	}

	// Generate swagger docs
	color.Blue("📚 Generating Swagger documentation...")
	swagCmd := exec.Command("swag", "init", "-g", "cmd/server/main.go", "-o", "docs/")
	swagCmd.Stdout = os.Stdout
	swagCmd.Stderr = os.Stderr
	
	if err := swagCmd.Run(); err != nil {
		return fmt.Errorf("failed to generate swagger docs: %w", err)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(swaggerCmd)
}