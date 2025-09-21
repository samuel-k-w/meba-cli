package cmd

import (
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	testWatch    bool
	testCoverage bool
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run tests",
	Long:  `Run unit tests for the meba application.`,
	Run: func(cmd *cobra.Command, args []string) {
		runTests()
	},
}

var e2eCmd = &cobra.Command{
	Use:   "e2e",
	Short: "Run end-to-end tests",
	Long:  `Run end-to-end tests for the meba application.`,
	Run: func(cmd *cobra.Command, args []string) {
		runE2ETests()
	},
}

func runTests() {
	color.Blue("🧪 Running tests...")
	
	args := []string{"test"}
	
	if testCoverage {
		args = append(args, "-coverprofile=coverage.out", "-covermode=atomic")
	}
	
	if testWatch {
		// For watch mode, we'd need to implement file watching
		// For now, just run tests once
		color.Yellow("⚠️  Watch mode not implemented yet, running tests once")
	}
	
	args = append(args, "./...")
	
	cmd := exec.Command("go", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		color.Red("Tests failed: %v", err)
		os.Exit(1)
	}

	if testCoverage {
		color.Green("✅ Tests completed with coverage!")
		color.Blue("📊 Generating coverage report...")
		
		// Generate HTML coverage report
		coverCmd := exec.Command("go", "tool", "cover", "-html=coverage.out", "-o=coverage.html")
		if err := coverCmd.Run(); err != nil {
			color.Yellow("⚠️  Could not generate HTML coverage report: %v", err)
		} else {
			color.Green("📄 Coverage report generated: coverage.html")
		}
	} else {
		color.Green("✅ Tests completed successfully!")
	}
}

func runE2ETests() {
	color.Blue("🧪 Running end-to-end tests...")
	
	// Check if test directory exists
	if _, err := os.Stat("test"); os.IsNotExist(err) {
		color.Yellow("⚠️  No test directory found, creating example e2e test structure...")
		createE2EStructure()
		return
	}
	
	args := []string{"test", "./test/..."}
	
	if testWatch {
		color.Yellow("⚠️  Watch mode not implemented yet, running tests once")
	}
	
	cmd := exec.Command("go", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		color.Red("E2E tests failed: %v", err)
		os.Exit(1)
	}

	color.Green("✅ E2E tests completed successfully!")
}

func createE2EStructure() {
	// Create test directory structure
	os.MkdirAll("test", 0755)
	
	// Create example e2e test
	exampleTest := `package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthEndpoint(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	
	// Add health endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Test
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "ok")
}

func TestAPIEndpoints(t *testing.T) {
	// Add your API endpoint tests here
	t.Skip("Implement your API tests")
}
`

	if err := os.WriteFile("test/e2e_test.go", []byte(exampleTest), 0644); err != nil {
		color.Red("Failed to create example e2e test: %v", err)
		return
	}
	
	color.Green("✅ Created example e2e test structure in test/ directory")
	color.Blue("💡 Run 'meba e2e' again to execute the tests")
}

func init() {
	rootCmd.AddCommand(testCmd)
	rootCmd.AddCommand(e2eCmd)
	
	testCmd.Flags().BoolVar(&testWatch, "watch", false, "Run tests in watch mode")
	testCmd.Flags().BoolVar(&testCoverage, "coverage", false, "Generate coverage report")
	
	e2eCmd.Flags().BoolVar(&testWatch, "watch", false, "Run e2e tests in watch mode")
}