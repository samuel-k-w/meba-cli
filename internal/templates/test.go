package templates

func HandlersTestGo() string {
	return `package internal

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandlers_SetupRoutes(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	router := gin.New()
	
	// Create handlers instance
	handlers := &Handlers{}
	
	// Setup routes
	handlers.SetupRoutes(router)
	
	// Test health endpoint
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/health", nil)
	router.ServeHTTP(w, req)
	
	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "ok")
}

func TestHealthEndpoint(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	router := gin.New()
	
	// Add health endpoint manually for testing
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
	assert.Contains(t, w.Body.String(), "Server is running")
}
`
}

func ServiceTestGo() string {
	return `package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	// Test service creation
	service := NewServices()
	
	// Assertions
	assert.NotNil(t, service)
	assert.IsType(t, &Services{}, service)
}

func TestService_SampleMethod(t *testing.T) {
	// Setup
	service := NewServices()
	
	// Test
	// Add your service method tests here
	
	// Assertions
	assert.NotNil(t, service)
	// Add more assertions based on your service methods
}

// Add more service tests as needed
func TestService_BusinessLogic(t *testing.T) {
	t.Skip("Implement your service business logic tests")
}
`
}