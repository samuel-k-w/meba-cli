package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	e2eWatchFlag bool
)

var e2eCmd = &cobra.Command{
	Use:   "e2e",
	Short: "Run end-to-end tests",
	Long:  "Run end-to-end tests from the test/ folder",
	Run: func(cmd *cobra.Command, args []string) {
		if e2eWatchFlag {
			runE2EWithWatch()
		} else {
			runE2E()
		}
	},
}

func init() {
	rootCmd.AddCommand(e2eCmd)
	e2eCmd.Flags().BoolVarP(&e2eWatchFlag, "watch", "w", false, "Run e2e tests in watch mode")
}

func runE2E() {
	fmt.Println("🔬 Running end-to-end tests...")
	
	// Check if test directory exists
	if _, err := os.Stat("test"); os.IsNotExist(err) {
		fmt.Println("📁 Creating test directory...")
		if err := os.MkdirAll("test", 0755); err != nil {
			fmt.Printf("❌ Failed to create test directory: %v\n", err)
			os.Exit(1)
		}
		
		// Create sample e2e test
		createSampleE2ETest()
		fmt.Println("📝 Created sample e2e test in test/ directory")
	}
	
	// Run tests in test directory
	cmd := exec.Command("go", "test", "./test/...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	if err := cmd.Run(); err != nil {
		fmt.Printf("❌ E2E tests failed: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Println("✅ All e2e tests passed!")
}

func runE2EWithWatch() {
	fmt.Println("👀 Running e2e tests in watch mode...")
	
	// Use gotestsum if available
	if _, err := exec.LookPath("gotestsum"); err == nil {
		cmd := exec.Command("gotestsum", "--watch", "./test/...")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		
		if err := cmd.Run(); err != nil {
			fmt.Printf("❌ E2E test watch failed: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("💡 Install gotestsum for better watch experience: go install gotest.tools/gotestsum@latest")
		runE2E()
	}
}

func createSampleE2ETest() {
	testContent := `package test

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
	router := gin.New()
	
	// Add health endpoint
	router.GET("/api/v1/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	// Test
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/health", nil)
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "ok")
}

func TestAPIEndpoints(t *testing.T) {
	// Add more e2e tests here
	t.Skip("Implement your e2e tests")
}
`
	
	testPath := filepath.Join("test", "e2e_test.go")
	if err := os.WriteFile(testPath, []byte(testContent), 0644); err != nil {
		fmt.Printf("Warning: Could not create sample e2e test: %v\n", err)
	}
}