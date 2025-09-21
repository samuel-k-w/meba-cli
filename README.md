# Meba CLI

A powerful CLI tool for generating Gin API projects inspired by NestJS architecture. Meba provides a complete development experience with dependency injection, hot reload, and modular structure.

## 🚀 Quick Install

### One-line install (Linux/macOS):
```bash
curl -fsSL https://raw.githubusercontent.com/meba-cli/meba/main/install.sh | bash
```

### Manual install:
```bash
go install github.com/meba-cli/meba/cmd/meba@latest
```

### From source:
```bash
git clone https://github.com/meba-cli/meba.git
cd meba
make install
```

## 📦 Features

- 🚀 **NestJS-inspired Architecture**: Modular structure with automatic dependency injection
- 🔥 **Hot Reload**: Development with Air for instant feedback
- 🏗️ **Dependency Injection**: Google Wire for clean, maintainable code
- 📊 **Database Integration**: GORM with PostgreSQL support
- 🔐 **Authentication**: JWT with Casbin RBAC/ABAC
- 📝 **Structured Logging**: Zap logger with different levels
- ✅ **Validation**: Request validation with go-playground/validator
- 📚 **API Documentation**: Swagger/OpenAPI auto-generation
- 🐳 **Docker Ready**: Complete containerization setup
- 🧪 **Testing**: Built-in testing utilities and e2e support

## 🎯 Quick Start

```bash
# Create a new project
meba new my-awesome-api
cd my-awesome-api

# Start development with hot reload
meba start --watch

# Generate a complete CRUD resource
meba g resource users

# Run tests
meba test
meba e2e

# Build for production
meba build
```

## 📚 Commands

### Project Creation
```bash
meba new <project-name>                    # Create new project
meba new <project-name> --skip-git         # Skip git initialization
meba new <project-name> --skip-install     # Skip go mod tidy
```

### Code Generation
```bash
meba g module <name>                       # Create module
meba g service <name>                      # Create service + test
meba g handler <name>                      # Create handler + test
meba g repository <name>                   # Create repository + test
meba g resource <name>                     # Complete CRUD resource
meba g middleware <name>                   # Create middleware
meba g guard <name>                        # Create guard

# Options
meba g service users --no-spec             # Skip test files
meba g handler users --dry-run             # Preview only
meba g module users --flat                 # Generate in current dir
```

### Build & Run
```bash
meba start                                 # Production mode
meba start --watch                         # Hot reload with Air
meba start --debug --watch                 # Debug mode
meba build                                 # Build to dist/
meba build --watch                         # Continuous compilation
```

### Testing
```bash
meba test                                  # Run unit tests
meba test --watch                          # Watch mode
meba test --coverage                       # Generate coverage
meba e2e                                   # End-to-end tests
meba e2e --watch                           # E2E watch mode
```

### Project Management
```bash
meba info                                  # Environment info
meba update                                # Update dependencies
meba version                               # Show version
```

## 🏗️ Generated Project Structure

```
my-awesome-api/
├── cmd/server/main.go                     # Application entry point
├── internal/
│   ├── app.go                            # Main app module registry
│   ├── users/                            # Example module
│   │   ├── module.go                     # Wire dependency set
│   │   ├── handlers.go + handlers_test.go
│   │   ├── service.go + service_test.go
│   │   ├── repository.go + repository_test.go
│   │   ├── entity.go                     # Domain models
│   │   └── dto.go                        # Request/response DTOs
│   └── wire.go + wire_gen.go             # Dependency injection
├── pkg/
│   ├── middleware/                       # Custom middleware
│   └── validator/                        # Validation utilities
├── test/e2e_test.go                      # End-to-end tests
├── configs/                              # Configuration files
├── dist/server                           # Compiled binary
└── docker-compose.yml                   # Docker setup
```

## 🔧 Development Workflow

1. **Create Project**: `meba new my-api`
2. **Generate Resources**: `meba g resource orders`
3. **Start Development**: `meba start --watch`
4. **Run Tests**: `meba test --watch`
5. **Build**: `meba build`

## 🤝 Contributing

1. Fork the repository
2. Create feature branch: `git checkout -b feature/amazing-feature`
3. Commit changes: `git commit -m 'Add amazing feature'`
4. Push to branch: `git push origin feature/amazing-feature`
5. Open Pull Request

## 📄 License

MIT License - see [LICENSE](LICENSE) file for details.

## 🆘 Support

- 📖 [Documentation](https://github.com/meba-cli/meba/wiki)
- 🐛 [Issues](https://github.com/meba-cli/meba/issues)
- 💬 [Discussions](https://github.com/meba-cli/meba/discussions)

---

**Happy coding with Meba! 🚀**