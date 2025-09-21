# Meba CLI

A powerful CLI tool for generating Gin API projects inspired by NestJS architecture. Meba provides a complete development experience with dependency injection, hot reload, and modular structure.

## Features

- 🚀 **NestJS-inspired Architecture**: Modular structure with automatic dependency injection
- 🔥 **Hot Reload**: Development with Air for instant feedback
- 🏗️ **Dependency Injection**: Google Wire for clean, maintainable code
- 📊 **Database Integration**: GORM with PostgreSQL support
- 🔐 **Authentication**: JWT with Casbin RBAC/ABAC
- 📝 **Structured Logging**: Zap logger with different levels
- ✅ **Validation**: Request validation with go-playground/validator
- 📚 **API Documentation**: Swagger/OpenAPI auto-generation
- 🐳 **Docker Ready**: Complete containerization setup
- 📈 **Monitoring**: Prometheus metrics and Grafana dashboards
- 🧪 **Testing**: Built-in testing utilities and e2e support

## Installation

```bash
go install github.com/meba-cli/meba@latest
```

## Quick Start

### Create a New Project

```bash
# Create a new meba project
meba new my-awesome-api

# Navigate to project
cd my-awesome-api

# Install dependencies
go mod tidy

# Start development server with hot reload
meba start --watch
```

## Commands Reference

### Project & Workspace

| Command | Description |
|---------|-------------|
| `meba new <project-name>` | Create a new meba application |
| `meba new <project-name> --skip-git` | Create project without Git initialization |
| `meba new <project-name> --directory <dir>` | Create project in specific directory |

### Code Generation

The `generate` (alias `g`) command creates boilerplate files with automatic dependency injection.

#### Available Schematics

| Schematic | Alias | Description |
|-----------|-------|-------------|
| `module` | - | Create a new module with complete structure |
| `handler` | `ha` | Create a handler (controller) |
| `service` | `s` | Create a service/provider |
| `repository` | `re` | Create a repository for data access |
| `resource` | - | Generate complete CRUD resource (module + handler + service + repository + DTOs) |
| `middleware` | - | Create a middleware |
| `guard` | - | Create an auth/route guard |
| `class` | - | Create a plain class file |
| `interface` | - | Create a Go interface |

#### Generation Examples

```bash
# Generate a complete module
meba g module users

# Generate individual components
meba g handler users
meba g service users  
meba g repository users

# Generate complete CRUD resource (recommended)
meba g resource products

# Generate middleware
meba g middleware auth
meba g guard admin

# Options
meba g service users --flat      # Generate in current directory
meba g handler users --no-spec   # Skip test files
meba g module users --dry-run    # Preview without creating files
```

### Running & Building

| Command | Description |
|---------|-------------|
| `meba start` | Start in production mode |
| `meba start --watch` | Start with hot reload (uses Air) |
| `meba start --debug --watch` | Debug mode with hot reload |
| `meba build` | Build the application |
| `meba build --watch` | Continuous compilation |

### Testing

| Command | Description |
|---------|-------------|
| `meba test` | Run unit tests |
| `meba test --watch` | Run tests in watch mode |
| `meba test --coverage` | Generate coverage report |
| `meba e2e` | Run end-to-end tests |
| `meba e2e --watch` | E2E tests in watch mode |

### Project Management

| Command | Description |
|---------|-------------|
| `meba info` | Show environment info and installed packages |
| `meba update` | Update packages to latest compatible versions |

## Project Structure

```
/my-awesome-api/
├── cmd/server/              # Application entry point
│   └── main.go
├── internal/                # Private application code
│   ├── app.go              # Main app module registry
│   ├── handlers.go         # Handler registry
│   ├── service.go          # Service registry  
│   ├── repository.go       # Repository registry
│   ├── entity.go           # Base entities
│   ├── dto.go              # Base DTOs
│   ├── wire.go             # Dependency injection
│   ├── wire_gen.go         # Generated DI code
│   ├── users/              # Example user module
│   │   ├── module.go       # Module wire set
│   │   ├── handlers.go     # HTTP handlers
│   │   ├── service.go      # Business logic
│   │   ├── repository.go   # Data access
│   │   ├── entity.go       # Domain models
│   │   └── dto.go          # Request/response DTOs
│   └── auth/               # Example auth module
│       └── ...
├── pkg/                    # Shared packages
│   ├── middleware/         # Custom middleware
│   └── validator/          # Validation utilities
├── configs/                # Configuration files
│   ├── config.yaml         # App configuration
│   └── config.go           # Config structs
├── deployments/            # Docker & deployment files
├── scripts/                # Build & utility scripts
├── go.mod
└── README.md
```

## Auto Dependency Injection

Meba automatically handles dependency injection using Google Wire:

### When you create a module:
- Automatically registers in top-level `app.go`
- Creates wire set for the module
- Sets up proper imports

### When you create components:
- Auto-injects into nearest module
- Updates module wire sets
- Maintains clean dependency graph

### Example Flow:

```bash
# 1. Create users module
meba g module users
# ✅ Creates internal/users/ with module.go
# ✅ Auto-registers in internal/app.go

# 2. Add service to users module  
meba g service users
# ✅ Creates internal/users/service.go
# ✅ Auto-adds to internal/users/module.go wire set

# 3. Add handler to users module
meba g handler users  
# ✅ Creates internal/users/handlers.go
# ✅ Auto-adds to internal/users/module.go wire set
# ✅ Ready to use with DI!
```

## Configuration

Meba uses Viper for flexible configuration supporting:
- YAML files
- Environment variables  
- Remote configuration
- Multiple config sources

Example `configs/config.yaml`:

```yaml
app:
  name: "my-api"
  port: 8080
  env: "development"

database:
  host: "localhost"
  port: 5432
  user: "postgres"
  password: "password"
  dbname: "myapi_db"

jwt:
  secret: "your-secret-key"
  expires_in: "24h"
```

## Development Workflow

### 1. Create New Feature Module

```bash
# Generate complete CRUD resource
meba g resource orders

# This creates:
# - internal/orders/module.go (wire set)
# - internal/orders/handlers.go (HTTP endpoints)  
# - internal/orders/service.go (business logic)
# - internal/orders/repository.go (data access)
# - internal/orders/entity.go (domain model)
# - internal/orders/dto.go (request/response types)
```

### 2. Start Development

```bash
# Start with hot reload
meba start --watch

# In another terminal, run tests
meba test --watch --coverage
```

### 3. Build & Deploy

```bash
# Build for production
meba build

# Or use Docker
docker-compose up --build
```

## Dependencies Included

- **Web Framework**: Gin
- **Dependency Injection**: Google Wire  
- **Configuration**: Viper
- **Database**: GORM + PostgreSQL
- **Authentication**: JWT + Casbin
- **Logging**: Zap
- **Validation**: go-playground/validator
- **JSON**: goccy/go-json (faster)
- **Cron Jobs**: robfig/cron
- **API Docs**: Swagger/swaggo
- **Testing**: testify + gomock
- **Monitoring**: Prometheus client
- **Hot Reload**: Air

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

- 📖 [Documentation](https://github.com/meba-cli/meba/wiki)
- 🐛 [Issue Tracker](https://github.com/meba-cli/meba/issues)
- 💬 [Discussions](https://github.com/meba-cli/meba/discussions)

---

**Happy coding with Meba! 🚀**