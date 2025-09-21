package templates

import "fmt"

func GoMod(projectName string) string {
	return fmt.Sprintf(`module %s

go 1.21

require (
	github.com/gin-gonic/gin v1.9.1
	github.com/google/wire v0.5.0
	github.com/spf13/viper v1.18.2
	gorm.io/gorm v1.25.5
	gorm.io/driver/sqlite v1.6.0
	gorm.io/driver/postgres v1.5.4
	github.com/golang-jwt/jwt/v5 v5.2.0
	github.com/casbin/casbin/v2 v2.81.0
	go.uber.org/zap v1.26.0
	github.com/go-playground/validator/v10 v10.16.0
	github.com/goccy/go-json v0.10.2
	github.com/robfig/cron/v3 v3.0.1
	github.com/swaggo/swag v1.16.2
	github.com/swaggo/gin-swagger v1.6.0
	github.com/swaggo/files v1.0.1
	github.com/stretchr/testify v1.8.4
	github.com/prometheus/client_golang v1.17.0
)
`, projectName)
}

func ReadmeMd(projectName string) string {
	return fmt.Sprintf(`# %s

A Gin API project inspired by NestJS architecture, generated with Meba CLI.

## Features

- 🚀 **NestJS-like Architecture**: Modular structure with dependency injection
- 🔥 **Hot Reload**: Development with Air
- 🏗️ **Dependency Injection**: Google Wire for clean DI
- 📊 **Database**: GORM with PostgreSQL
- 🔐 **Authentication**: JWT with Casbin RBAC
- 📝 **Logging**: Structured logging with Zap
- ✅ **Validation**: Request validation with go-playground/validator
- 📚 **API Docs**: Swagger/OpenAPI documentation
- 🐳 **Docker**: Ready for containerization
- 📈 **Monitoring**: Prometheus metrics

## Getting Started

### Prerequisites

- Go 1.21+
- PostgreSQL
- Air (for hot reload): ` + "`go install github.com/cosmtrek/air@latest`" + `

### Installation

1. Clone the repository
2. Install dependencies:
   ` + "```bash" + `
   go mod tidy
   ` + "```" + `

3. Copy and configure environment:
   ` + "```bash" + `
   cp configs/config.yaml.example configs/config.yaml
   ` + "```" + `

4. Run the application:
   ` + "```bash" + `
   # Development with hot reload
   meba start --watch
   
   # Production
   meba start
   ` + "```" + `

## Project Structure

` + "```" + `
/%s/
├── cmd/server/          # Application entry point
├── internal/            # Private application code
│   ├── users/          # Example user module
│   ├── auth/           # Example auth module
│   └── ...
├── pkg/                # Shared packages
├── configs/            # Configuration files
├── deployments/        # Docker and deployment files
└── scripts/           # Build and utility scripts
` + "```" + `

## Development

### Generate New Module

` + "```bash" + `
meba g module products
meba g service products
meba g handler products
meba g repository products
` + "```" + `

### Generate Complete Resource

` + "```bash" + `
meba g resource orders
` + "```" + `

### Testing

` + "```bash" + `
go test ./...
` + "```" + `

## API Documentation

Visit ` + "`http://localhost:8080/swagger/index.html`" + ` for API documentation.

## License

MIT License
`, projectName, projectName)
}

func MainGo(projectName string) string {
	return fmt.Sprintf(`package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"%s/internal"
	"%s/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
)

// @title %s API
// @version 1.0
// @description A Gin API project inspired by NestJS architecture
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1
// @schemes http https

func main() {
	// Initialize logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Initialize Gin
	r := gin.Default()

	// Initialize application
	app, cleanup, err := internal.InitializeApp()
	if err != nil {
		log.Fatal("Failed to initialize app:", err)
	}
	defer cleanup()

	// Setup Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	// Setup routes
	app.SetupRoutes(r)

	// Start server
	logger.Info("Starting server on :8080")
	logger.Info("Swagger docs available at: http://localhost:8080/swagger/index.html")
	if err := r.Run(":8080"); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}`, projectName, projectName, projectName)
}

func AppGo() string {
	return `package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// App represents the main application
type App struct {
	handlers *Handlers
}

// NewApp creates a new application instance
func NewApp(handlers *Handlers) *App {
	return &App{
		handlers: handlers,
	}
}

// SetupRoutes configures all application routes
func (a *App) SetupRoutes(r *gin.Engine) {
	a.handlers.SetupRoutes(r)
}

// AppSet is the wire set for the main application
var AppSet = wire.NewSet(
	NewApp,
	HandlersSet,
	ServiceSet,
	RepositorySet,
)
`
}

func HandlersGo() string {
	return `package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// Handlers aggregates all handler modules
type Handlers struct {
	// Add your module handlers here
	// users *users.Handlers
}

// NewHandlers creates a new handlers instance
func NewHandlers() *Handlers {
	return &Handlers{}
}

// SetupRoutes configures all routes
func (h *Handlers) SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	
	// Health check
	// @Summary Health Check
	// @Description Check if the API is running
	// @Tags health
	// @Accept json
	// @Produce json
	// @Success 200 {object} map[string]string
	// @Router /health [get]
	api.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "message": "Server is running"})
	})

	// Setup module routes here
	// h.users.SetupRoutes(api)
}

// HandlersSet is the wire set for handlers
var HandlersSet = wire.NewSet(
	NewHandlers,
)
`
}

func ServiceGo() string {
	return `package internal

import (
	"github.com/google/wire"
)

// Services aggregates all service modules
type Services struct {
	// Add your module services here
	// Users *users.Service
}

// NewServices creates a new services instance
func NewServices() *Services {
	return &Services{}
}

// ServiceSet is the wire set for services
var ServiceSet = wire.NewSet(
	NewServices,
)
`
}

func EntityGo() string {
	return `package internal

import (
	"time"
	"gorm.io/gorm"
)

// BaseEntity provides common fields for all entities
type BaseEntity struct {
	ID        uint           ` + "`json:\"id\" gorm:\"primarykey\"`" + `
	CreatedAt time.Time      ` + "`json:\"created_at\"`" + `
	UpdatedAt time.Time      ` + "`json:\"updated_at\"`" + `
	DeletedAt gorm.DeletedAt ` + "`json:\"deleted_at\" gorm:\"index\"`" + `
}
`
}

func DtoGo() string {
	return `package internal

// BaseResponse provides common response structure
type BaseResponse struct {
	Success bool        ` + "`json:\"success\"`" + `
	Message string      ` + "`json:\"message\"`" + `
	Data    interface{} ` + "`json:\"data,omitempty\"`" + `
	Error   string      ` + "`json:\"error,omitempty\"`" + `
}

// PaginationRequest provides common pagination parameters
type PaginationRequest struct {
	Page     int ` + "`json:\"page\" form:\"page\" validate:\"min=1\"`" + `
	PageSize int ` + "`json:\"page_size\" form:\"page_size\" validate:\"min=1,max=100\"`" + `
}

// PaginationResponse provides common pagination response
type PaginationResponse struct {
	Page       int         ` + "`json:\"page\"`" + `
	PageSize   int         ` + "`json:\"page_size\"`" + `
	Total      int64       ` + "`json:\"total\"`" + `
	TotalPages int         ` + "`json:\"total_pages\"`" + `
	Data       interface{} ` + "`json:\"data\"`" + `
}
`
}

func RepositoryGo() string {
	return `package internal

import (
	"gorm.io/gorm"
	"github.com/google/wire"
)

// Repositories aggregates all repository modules
type Repositories struct {
	db *gorm.DB
	// Add your module repositories here
	// Users *users.Repository
}

// NewRepositories creates a new repositories instance
func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		db: db,
	}
}

// RepositorySet is the wire set for repositories
var RepositorySet = wire.NewSet(
	NewRepositories,
	NewDatabase,
)

// NewDatabase creates a new database connection
func NewDatabase() (*gorm.DB, error) {
	// This should be implemented with actual database configuration
	// For now, return nil - implement based on your database choice
	return nil, nil
}
`
}

func WireGo() string {
	return `//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/google/wire"
)

// InitializeApp initializes the application with dependency injection
func InitializeApp() (*App, func(), error) {
	wire.Build(AppSet)
	return nil, nil, nil
}
`
}