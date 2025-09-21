# Meba CLI Usage Guide

## Complete CLI Application for Gin API Generation

This is a fully functional CLI tool that generates Gin API projects inspired by NestJS architecture with automatic dependency injection using Google Wire.

## ✅ What's Implemented

### 1. **Project Generation**
```bash
meba new my-api                    # Create new project
meba new my-api --skip-git         # Skip git initialization
meba new my-api --directory ./apps # Custom directory
```

### 2. **Code Scaffolding with Auto-Injection**
```bash
# Generate complete module with all components
meba g module users
meba g resource products           # Complete CRUD resource

# Generate individual components
meba g handler users
meba g service users
meba g repository users

# Generate utilities
meba g middleware cors
meba g guard admin
```

### 3. **Development Commands**
```bash
meba start --watch                 # Hot reload with Air
meba start --debug --watch         # Debug mode
meba build                         # Production build
meba test --coverage               # Run tests with coverage
meba e2e                          # End-to-end tests
```

### 4. **Project Management**
```bash
meba info                         # Environment information
meba update                       # Update dependencies
```

## 🏗️ Generated Project Structure

```
/my-api/
├── cmd/server/main.go           # Application entry point
├── internal/                    # Private application code
│   ├── app.go                  # Main app module registry
│   ├── handlers.go             # Handler registry
│   ├── service.go              # Service registry
│   ├── repository.go           # Repository registry
│   ├── entity.go               # Base entities
│   ├── dto.go                  # Base DTOs
│   ├── wire.go                 # Dependency injection
│   └── users/                  # Generated module
│       ├── module.go           # Wire set
│       ├── handlers.go         # HTTP handlers
│       ├── service.go          # Business logic
│       ├── repository.go       # Data access
│       ├── entity.go           # Domain models
│       └── dto.go              # Request/response DTOs
├── pkg/                        # Shared packages
│   ├── middleware/             # Custom middleware
│   └── validator/              # Validation utilities
├── configs/                    # Configuration files
├── deployments/                # Docker & deployment
├── .air.toml                   # Hot reload config
├── docker-compose.yml          # Multi-service setup
└── Dockerfile                  # Container definition
```

## 🔄 Auto Dependency Injection

### When you create a module:
1. ✅ Creates complete module structure
2. ✅ Auto-registers in `internal/app.go`
3. ✅ Sets up Wire dependency injection
4. ✅ Updates imports automatically

### When you create components:
1. ✅ Auto-injects into nearest module
2. ✅ Updates module wire sets
3. ✅ Maintains clean dependency graph

## 📦 Included Dependencies

- **Web Framework**: Gin
- **Dependency Injection**: Google Wire
- **Configuration**: Viper (YAML, ENV, remote)
- **Database**: GORM + PostgreSQL
- **Authentication**: JWT + Casbin RBAC
- **Logging**: Zap structured logging
- **Validation**: go-playground/validator
- **JSON**: goccy/go-json (faster)
- **Cron Jobs**: robfig/cron
- **API Docs**: Swagger/swaggo
- **Testing**: testify + gomock
- **Monitoring**: Prometheus client
- **Hot Reload**: Air

## 🚀 Example Workflow

```bash
# 1. Create new project
meba new ecommerce-api
cd ecommerce-api

# 2. Generate complete CRUD resources
meba g resource users
meba g resource products
meba g resource orders

# 3. Generate custom middleware
meba g middleware rate-limit
meba g guard admin

# 4. Start development with hot reload
meba start --watch

# 5. Run tests
meba test --coverage

# 6. Build for production
meba build
```

## 🔧 Build & Install

```bash
# Build locally
make build

# Install globally
make install

# Run development setup
make dev-setup

# Build for multiple platforms
make release
```

## 🐳 Docker Support

Generated projects include:
- ✅ Dockerfile for containerization
- ✅ docker-compose.yml with PostgreSQL, Redis
- ✅ Prometheus & Grafana monitoring
- ✅ Health checks and proper networking

## 🧪 Testing Support

- ✅ Unit test scaffolding
- ✅ E2E test structure
- ✅ Coverage reporting
- ✅ Watch mode for continuous testing

## 📊 Monitoring & Observability

- ✅ Prometheus metrics endpoint
- ✅ Structured logging with Zap
- ✅ Health check endpoints
- ✅ Grafana dashboard ready

## 🎯 Key Features Achieved

1. **✅ NestJS-like Architecture**: Complete modular structure
2. **✅ Automatic DI**: Google Wire integration with auto-registration
3. **✅ Hot Reload**: Air integration for development
4. **✅ Complete CRUD**: Full resource generation
5. **✅ Production Ready**: Docker, monitoring, logging
6. **✅ Developer Experience**: Rich CLI with helpful commands
7. **✅ Extensible**: Easy to add new generators and templates

This is a complete, production-ready CLI tool that generates modern Go APIs with enterprise-grade architecture and developer experience comparable to NestJS.