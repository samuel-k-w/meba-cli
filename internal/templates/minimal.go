package templates

import (
	"fmt"
	"strings"
)

// MinimalModuleGo creates just a module.go file with empty wire set
func MinimalModuleGo(name string) string {
	return fmt.Sprintf(`package %s

import (
	"github.com/google/wire"
)

// Module is the wire set for the %s module
var Module = wire.NewSet()
`, name, name)
}

// MinimalServiceGo creates just a service struct with constructor
func MinimalServiceGo(name string) string {
	titleName := strings.Title(name)
	return fmt.Sprintf(`package %s

// Service handles business logic for %s
type Service struct {
	// Add dependencies here
}

// New%sService creates a new service instance
func New%sService() *Service {
	return &Service{}
}
`, name, name, titleName, titleName)
}

// MinimalRepositoryGo creates just a repository struct with constructor
func MinimalRepositoryGo(name string) string {
	titleName := strings.Title(name)
	return fmt.Sprintf(`package %s

import (
	"gorm.io/gorm"
)

// Repository handles data access for %s
type Repository struct {
	db *gorm.DB
}

// New%sRepository creates a new repository instance
func New%sRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}
`, name, name, titleName, titleName)
}

// MinimalHandlersGo creates handlers with SetupRoutes but no route methods
func MinimalHandlersGo(name string) string {
	titleName := strings.Title(name)
	return fmt.Sprintf(`package %s

import (
	"github.com/gin-gonic/gin"
)

// Handlers handles HTTP requests for %s
type Handlers struct {
	service *Service
}

// New%sHandlers creates a new handlers instance
func New%sHandlers(service *Service) *Handlers {
	return &Handlers{
		service: service,
	}
}

// SetupRoutes configures routes for %s module
func (h *Handlers) SetupRoutes(r *gin.RouterGroup) {
	// %sGroup := r.Group("/%s")
	// {
	//     Add your routes here
	//     %sGroup.GET("", h.GetAll)
	//     %sGroup.POST("", h.Create)
	// }
}
`, name, name, titleName, titleName, name, name, name, name, name)
}