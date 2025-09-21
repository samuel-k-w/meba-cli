package templates

import (
	"fmt"
	"strings"
)

// ServiceTestGoModule creates test file for a specific service module
func ServiceTestGoModule(name string) string {
	titleName := strings.Title(name)
	return fmt.Sprintf(`package %s

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew%sService(t *testing.T) {
	// Test service creation
	service := New%sService()
	
	// Assertions
	assert.NotNil(t, service)
	assert.IsType(t, &Service{}, service)
}

func Test%sService_BusinessLogic(t *testing.T) {
	// Setup
	service := New%sService()
	
	// Test your service methods here
	// Example:
	// result := service.SomeMethod()
	// assert.NotNil(t, result)
	
	// Assertions
	assert.NotNil(t, service)
}

// Add more %s service tests as needed
func Test%sService_Methods(t *testing.T) {
	t.Skip("Implement your %s service method tests")
}
`, name, titleName, titleName, titleName, titleName, name, titleName, name)
}

// RepositoryTestGoModule creates test file for a specific repository module
func RepositoryTestGoModule(name string) string {
	titleName := strings.Title(name)
	return fmt.Sprintf(`package %s

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestNew%sRepository(t *testing.T) {
	// Setup test database
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)
	
	// Test repository creation
	repo := New%sRepository(db)
	
	// Assertions
	assert.NotNil(t, repo)
	assert.IsType(t, &Repository{}, repo)
	assert.Equal(t, db, repo.db)
}

func Test%sRepository_DatabaseOperations(t *testing.T) {
	// Setup test database
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)
	
	// Auto-migrate the schema
	err = db.AutoMigrate(&%s{})
	assert.NoError(t, err)
	
	// Create repository
	repo := New%sRepository(db)
	
	// Test your repository methods here
	// Example:
	// entity := &%s{Name: "test"}
	// err = repo.Create(entity)
	// assert.NoError(t, err)
	
	assert.NotNil(t, repo)
}

// Add more %s repository tests as needed
func Test%sRepository_CRUD(t *testing.T) {
	t.Skip("Implement your %s repository CRUD tests")
}
`, name, titleName, titleName, titleName, titleName, titleName, titleName, name, titleName, name)
}

// HandlersTestGoModule creates test file for a specific handlers module
func HandlersTestGoModule(name string) string {
	titleName := strings.Title(name)
	return fmt.Sprintf(`package %s

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNew%sHandlers(t *testing.T) {
	// Setup
	service := &Service{} // Mock service
	
	// Test handlers creation
	handlers := New%sHandlers(service)
	
	// Assertions
	assert.NotNil(t, handlers)
	assert.IsType(t, &Handlers{}, handlers)
	assert.Equal(t, service, handlers.service)
}

func Test%sHandlers_SetupRoutes(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	router := gin.New()
	service := &Service{} // Mock service
	handlers := New%sHandlers(service)
	
	// Setup routes
	apiGroup := router.Group("/api/v1")
	handlers.SetupRoutes(apiGroup)
	
	// Test that routes are registered
	routes := router.Routes()
	assert.NotEmpty(t, routes)
}

func Test%sHandlers_HTTPEndpoints(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	router := gin.New()
	service := &Service{} // Mock service
	handlers := New%sHandlers(service)
	
	// Setup routes
	apiGroup := router.Group("/api/v1")
	handlers.SetupRoutes(apiGroup)
	
	// Test endpoints (add specific tests for your handlers)
	// Example:
	// w := httptest.NewRecorder()
	// req, _ := http.NewRequest("GET", "/api/v1/%s", nil)
	// router.ServeHTTP(w, req)
	// assert.Equal(t, http.StatusOK, w.Code)
	
	assert.NotNil(t, handlers)
}

// Add more %s handlers tests as needed
func Test%sHandlers_Methods(t *testing.T) {
	t.Skip("Implement your %s handlers method tests")
}
`, name, titleName, titleName, titleName, titleName, titleName, titleName, name, name, titleName, name)
}