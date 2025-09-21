package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/meba-cli/meba/internal/generator"
	"github.com/spf13/cobra"
)

var (
	dryRun  bool
	flat    bool
	noSpec  bool
	project string
)

var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"g"},
	Short:   "Generate code scaffolding",
	Long:    `Generate various code scaffolding like modules, services, handlers, repositories, etc.`,
}

var moduleCmd = &cobra.Command{
	Use:   "module [name]",
	Short: "Generate a new module",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		if err := generator.GenerateModule(name, dryRun, flat); err != nil {
			color.Red("Error generating module: %v", err)
			os.Exit(1)
		}
		color.Green("✅ Module '%s' generated successfully!", name)
	},
}

var handlerCmd = &cobra.Command{
	Use:     "handler [name]",
	Aliases: []string{"ha"},
	Short:   "Generate a new handler",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		if err := generator.GenerateHandler(name, dryRun, flat, noSpec); err != nil {
			color.Red("Error generating handler: %v", err)
			os.Exit(1)
		}
		color.Green("✅ Handler '%s' generated successfully!", name)
	},
}

var serviceCmd = &cobra.Command{
	Use:     "service [name]",
	Aliases: []string{"s"},
	Short:   "Generate a new service",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		if err := generator.GenerateService(name, dryRun, flat, noSpec); err != nil {
			color.Red("Error generating service: %v", err)
			os.Exit(1)
		}
		color.Green("✅ Service '%s' generated successfully!", name)
	},
}

var repositoryCmd = &cobra.Command{
	Use:     "repository [name]",
	Aliases: []string{"re"},
	Short:   "Generate a new repository",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		if err := generator.GenerateRepository(name, dryRun, flat, noSpec); err != nil {
			color.Red("Error generating repository: %v", err)
			os.Exit(1)
		}
		color.Green("✅ Repository '%s' generated successfully!", name)
	},
}

var resourceCmd = &cobra.Command{
	Use:   "resource [name]",
	Short: "Generate a complete CRUD resource",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		if err := generator.GenerateResource(name, dryRun, noSpec); err != nil {
			color.Red("Error generating resource: %v", err)
			os.Exit(1)
		}
		color.Green("✅ Resource '%s' generated successfully!", name)
	},
}

var middlewareCmd = &cobra.Command{
	Use:   "middleware [name]",
	Short: "Generate a new middleware",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		if err := generator.GenerateMiddleware(name, dryRun, flat); err != nil {
			color.Red("Error generating middleware: %v", err)
			os.Exit(1)
		}
		color.Green("✅ Middleware '%s' generated successfully!", name)
	},
}

var guardCmd = &cobra.Command{
	Use:   "guard [name]",
	Short: "Generate a new guard",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		if err := generator.GenerateGuard(name, dryRun, flat); err != nil {
			color.Red("Error generating guard: %v", err)
			os.Exit(1)
		}
		color.Green("✅ Guard '%s' generated successfully!", name)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	
	// Add subcommands
	generateCmd.AddCommand(moduleCmd)
	generateCmd.AddCommand(handlerCmd)
	generateCmd.AddCommand(serviceCmd)
	generateCmd.AddCommand(repositoryCmd)
	generateCmd.AddCommand(resourceCmd)
	generateCmd.AddCommand(middlewareCmd)
	generateCmd.AddCommand(guardCmd)

	// Add flags to all generate commands
	for _, cmd := range []*cobra.Command{moduleCmd, handlerCmd, serviceCmd, repositoryCmd, resourceCmd, middlewareCmd, guardCmd} {
		cmd.Flags().BoolVar(&dryRun, "dry-run", false, "Show files to be created without writing")
		cmd.Flags().BoolVar(&flat, "flat", false, "Generate files in current directory")
		cmd.Flags().BoolVar(&noSpec, "no-spec", false, "Skip test files")
		cmd.Flags().StringVar(&project, "project", "", "Project name for monorepo")
	}
}