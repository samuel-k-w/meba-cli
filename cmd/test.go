package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	testWatchFlag    bool
	testCoverageFlag bool
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run unit tests",
	Long:  "Run unit tests with optional watch mode and coverage",
	Run: func(cmd *cobra.Command, args []string) {
		if testWatchFlag {
			runTestsWithWatch()
		} else if testCoverageFlag {
			runTestsWithCoverage()
		} else {
			runTests()
		}
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.Flags().BoolVarP(&testWatchFlag, "watch", "w", false, "Run tests in watch mode")
	testCmd.Flags().BoolVar(&testCoverageFlag, "coverage", false, "Generate code coverage report")
}

func runTests() {
	fmt.Println("🧪 Running unit tests...")
	
	cmd := exec.Command("go", "test", "./...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	if err := cmd.Run(); err != nil {
		fmt.Printf("❌ Tests failed: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Println("✅ All tests passed!")
}

func runTestsWithWatch() {
	fmt.Println("👀 Running tests in watch mode...")
	
	// Use gotestsum if available, otherwise fall back to basic watch
	if _, err := exec.LookPath("gotestsum"); err == nil {
		cmd := exec.Command("gotestsum", "--watch", "./...")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		
		if err := cmd.Run(); err != nil {
			fmt.Printf("❌ Test watch failed: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("💡 Install gotestsum for better watch experience: go install gotest.tools/gotestsum@latest")
		// Simple watch implementation
		runTests()
	}
}

func runTestsWithCoverage() {
	fmt.Println("📊 Running tests with coverage...")
	
	// Create coverage directory
	if err := os.MkdirAll("coverage", 0755); err != nil {
		fmt.Printf("❌ Failed to create coverage directory: %v\n", err)
		os.Exit(1)
	}
	
	// Run tests with coverage
	cmd := exec.Command("go", "test", "-coverprofile=coverage/coverage.out", "./...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	if err := cmd.Run(); err != nil {
		fmt.Printf("❌ Tests with coverage failed: %v\n", err)
		os.Exit(1)
	}
	
	// Generate HTML coverage report
	cmd = exec.Command("go", "tool", "cover", "-html=coverage/coverage.out", "-o", "coverage/coverage.html")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Warning: Could not generate HTML coverage report: %v\n", err)
	} else {
		fmt.Println("📈 Coverage report generated: coverage/coverage.html")
	}
	
	// Show coverage summary
	cmd = exec.Command("go", "tool", "cover", "-func=coverage/coverage.out")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	
	fmt.Println("✅ Tests completed with coverage!")
}