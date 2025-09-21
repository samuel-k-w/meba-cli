package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/meba-cli/meba/internal/generator"
	"github.com/spf13/cobra"
)

var (
	skipGit     bool
	skipInstall bool
	directory   string
)

var newCmd = &cobra.Command{
	Use:   "new [project-name]",
	Short: "Create a new meba application",
	Long:  `Create a new meba application with the complete project structure and dependencies.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		
		targetDir := directory
		if targetDir == "" {
			targetDir = projectName
		}

		if err := generator.CreateProject(projectName, targetDir, skipGit, skipInstall); err != nil {
			color.Red("Error creating project: %v", err)
			os.Exit(1)
		}

		color.Green("✅ Project '%s' created successfully!", projectName)
		fmt.Printf("\nNext steps:\n")
		fmt.Printf("  cd %s\n", targetDir)
		fmt.Printf("  go mod tidy\n")
		fmt.Printf("  meba start --watch\n")
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().BoolVar(&skipGit, "skip-git", false, "Skip git repository initialization")
	newCmd.Flags().BoolVar(&skipInstall, "skip-install", false, "Skip go mod tidy during project creation")
	newCmd.Flags().StringVar(&directory, "directory", "", "Specify directory name for the project")
}