package templates

import "fmt"

func ConfigYaml() string {
	return `# Application Configuration
app:
  name: "meba-app"
  version: "1.0.0"
  port: 8080
  env: "development"
  debug: true

# Database Configuration
database:
  host: "localhost"
  port: 5432
  user: "postgres"
  password: "password"
  dbname: "meba_db"
  sslmode: "disable"
  timezone: "UTC"
  max_idle_conns: 10
  max_open_conns: 100
  conn_max_lifetime: "1h"

# Redis Configuration
redis:
  host: "localhost"
  port: 6379
  password: ""
  db: 0
  pool_size: 10

# JWT Configuration
jwt:
  secret: "your-super-secret-jwt-key-change-this-in-production"
  expires_in: "24h"
  refresh_expires_in: "168h"

# CORS Configuration
cors:
  allowed_origins: ["*"]
  allowed_methods: ["GET", "POST", "PUT", "DELETE", "OPTIONS"]
  allowed_headers: ["*"]
  allow_credentials: true

# Rate Limiting
rate_limit:
  enabled: true
  requests_per_minute: 100
  burst: 10

# Logging Configuration
logging:
  level: "info"
  format: "json"
  output: "stdout"

# Swagger Configuration
swagger:
  enabled: true
  title: "Meba API"
  description: "A Gin API inspired by NestJS"
  version: "1.0.0"
  host: "localhost:8080"
  base_path: "/api/v1"

# Monitoring
monitoring:
  prometheus:
    enabled: true
    path: "/metrics"
  health_check:
    enabled: true
    path: "/health"

# File Upload
upload:
  max_file_size: "10MB"
  allowed_types: ["image/jpeg", "image/png", "image/gif", "application/pdf"]
  upload_path: "./uploads"

# Email Configuration
email:
  smtp_host: "smtp.gmail.com"
  smtp_port: 587
  username: "your-email@gmail.com"
  password: "your-app-password"
  from_name: "Meba App"
  from_email: "noreply@meba.com"

# Cache Configuration
cache:
  default_ttl: "1h"
  cleanup_interval: "10m"
`
}

func ConfigGo() string {
	return `package configs

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Config holds all configuration for our application
type Config struct {
	App        AppConfig        ` + "`mapstructure:\"app\"`" + `
	Database   DatabaseConfig   ` + "`mapstructure:\"database\"`" + `
	Redis      RedisConfig      ` + "`mapstructure:\"redis\"`" + `
	JWT        JWTConfig        ` + "`mapstructure:\"jwt\"`" + `
	CORS       CORSConfig       ` + "`mapstructure:\"cors\"`" + `
	RateLimit  RateLimitConfig  ` + "`mapstructure:\"rate_limit\"`" + `
	Logging    LoggingConfig    ` + "`mapstructure:\"logging\"`" + `
	Swagger    SwaggerConfig    ` + "`mapstructure:\"swagger\"`" + `
	Monitoring MonitoringConfig ` + "`mapstructure:\"monitoring\"`" + `
	Upload     UploadConfig     ` + "`mapstructure:\"upload\"`" + `
	Email      EmailConfig      ` + "`mapstructure:\"email\"`" + `
	Cache      CacheConfig      ` + "`mapstructure:\"cache\"`" + `
}

type AppConfig struct {
	Name    string ` + "`mapstructure:\"name\"`" + `
	Version string ` + "`mapstructure:\"version\"`" + `
	Port    int    ` + "`mapstructure:\"port\"`" + `
	Env     string ` + "`mapstructure:\"env\"`" + `
	Debug   bool   ` + "`mapstructure:\"debug\"`" + `
}

type DatabaseConfig struct {
	Host            string        ` + "`mapstructure:\"host\"`" + `
	Port            int           ` + "`mapstructure:\"port\"`" + `
	User            string        ` + "`mapstructure:\"user\"`" + `
	Password        string        ` + "`mapstructure:\"password\"`" + `
	DBName          string        ` + "`mapstructure:\"dbname\"`" + `
	SSLMode         string        ` + "`mapstructure:\"sslmode\"`" + `
	Timezone        string        ` + "`mapstructure:\"timezone\"`" + `
	MaxIdleConns    int           ` + "`mapstructure:\"max_idle_conns\"`" + `
	MaxOpenConns    int           ` + "`mapstructure:\"max_open_conns\"`" + `
	ConnMaxLifetime time.Duration ` + "`mapstructure:\"conn_max_lifetime\"`" + `
}

type RedisConfig struct {
	Host     string ` + "`mapstructure:\"host\"`" + `
	Port     int    ` + "`mapstructure:\"port\"`" + `
	Password string ` + "`mapstructure:\"password\"`" + `
	DB       int    ` + "`mapstructure:\"db\"`" + `
	PoolSize int    ` + "`mapstructure:\"pool_size\"`" + `
}

type JWTConfig struct {
	Secret           string        ` + "`mapstructure:\"secret\"`" + `
	ExpiresIn        time.Duration ` + "`mapstructure:\"expires_in\"`" + `
	RefreshExpiresIn time.Duration ` + "`mapstructure:\"refresh_expires_in\"`" + `
}

type CORSConfig struct {
	AllowedOrigins   []string ` + "`mapstructure:\"allowed_origins\"`" + `
	AllowedMethods   []string ` + "`mapstructure:\"allowed_methods\"`" + `
	AllowedHeaders   []string ` + "`mapstructure:\"allowed_headers\"`" + `
	AllowCredentials bool     ` + "`mapstructure:\"allow_credentials\"`" + `
}

type RateLimitConfig struct {
	Enabled           bool ` + "`mapstructure:\"enabled\"`" + `
	RequestsPerMinute int  ` + "`mapstructure:\"requests_per_minute\"`" + `
	Burst             int  ` + "`mapstructure:\"burst\"`" + `
}

type LoggingConfig struct {
	Level  string ` + "`mapstructure:\"level\"`" + `
	Format string ` + "`mapstructure:\"format\"`" + `
	Output string ` + "`mapstructure:\"output\"`" + `
}

type SwaggerConfig struct {
	Enabled     bool   ` + "`mapstructure:\"enabled\"`" + `
	Title       string ` + "`mapstructure:\"title\"`" + `
	Description string ` + "`mapstructure:\"description\"`" + `
	Version     string ` + "`mapstructure:\"version\"`" + `
	Host        string ` + "`mapstructure:\"host\"`" + `
	BasePath    string ` + "`mapstructure:\"base_path\"`" + `
}

type MonitoringConfig struct {
	Prometheus  PrometheusConfig  ` + "`mapstructure:\"prometheus\"`" + `
	HealthCheck HealthCheckConfig ` + "`mapstructure:\"health_check\"`" + `
}

type PrometheusConfig struct {
	Enabled bool   ` + "`mapstructure:\"enabled\"`" + `
	Path    string ` + "`mapstructure:\"path\"`" + `
}

type HealthCheckConfig struct {
	Enabled bool   ` + "`mapstructure:\"enabled\"`" + `
	Path    string ` + "`mapstructure:\"path\"`" + `
}

type UploadConfig struct {
	MaxFileSize  string   ` + "`mapstructure:\"max_file_size\"`" + `
	AllowedTypes []string ` + "`mapstructure:\"allowed_types\"`" + `
	UploadPath   string   ` + "`mapstructure:\"upload_path\"`" + `
}

type EmailConfig struct {
	SMTPHost  string ` + "`mapstructure:\"smtp_host\"`" + `
	SMTPPort  int    ` + "`mapstructure:\"smtp_port\"`" + `
	Username  string ` + "`mapstructure:\"username\"`" + `
	Password  string ` + "`mapstructure:\"password\"`" + `
	FromName  string ` + "`mapstructure:\"from_name\"`" + `
	FromEmail string ` + "`mapstructure:\"from_email\"`" + `
}

type CacheConfig struct {
	DefaultTTL      time.Duration ` + "`mapstructure:\"default_ttl\"`" + `
	CleanupInterval time.Duration ` + "`mapstructure:\"cleanup_interval\"`" + `
}

// LoadConfig loads configuration from file and environment variables
func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// Enable environment variable support
	viper.AutomaticEnv()

	// Set default values
	setDefaults()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}

func setDefaults() {
	viper.SetDefault("app.name", "meba-app")
	viper.SetDefault("app.version", "1.0.0")
	viper.SetDefault("app.port", 8080)
	viper.SetDefault("app.env", "development")
	viper.SetDefault("app.debug", true)

	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 5432)
	viper.SetDefault("database.user", "postgres")
	viper.SetDefault("database.password", "password")
	viper.SetDefault("database.dbname", "meba_db")
	viper.SetDefault("database.sslmode", "disable")
	viper.SetDefault("database.timezone", "UTC")
	viper.SetDefault("database.max_idle_conns", 10)
	viper.SetDefault("database.max_open_conns", 100)
	viper.SetDefault("database.conn_max_lifetime", "1h")

	viper.SetDefault("jwt.secret", "change-this-in-production")
	viper.SetDefault("jwt.expires_in", "24h")
	viper.SetDefault("jwt.refresh_expires_in", "168h")

	viper.SetDefault("logging.level", "info")
	viper.SetDefault("logging.format", "json")
	viper.SetDefault("logging.output", "stdout")
}

// GetDSN returns the database connection string
func (d *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		d.Host, d.Port, d.User, d.Password, d.DBName, d.SSLMode, d.Timezone)
}

// IsProduction returns true if the app is running in production
func (a *AppConfig) IsProduction() bool {
	return a.Env == "production"
}

// IsDevelopment returns true if the app is running in development
func (a *AppConfig) IsDevelopment() bool {
	return a.Env == "development"
}
`
}

func AirToml() string {
	return `root = "."
testdata_dir = "testdata"
tmp_dir = "tmp"

[build]
  args_bin = []
  bin = "./tmp/main"
  cmd = "go build -o ./tmp/main ./cmd/server"
  delay = 1000
  exclude_dir = ["assets", "tmp", "vendor", "testdata"]
  exclude_file = []
  exclude_regex = ["_test.go"]
  exclude_unchanged = false
  follow_symlink = false
  full_bin = ""
  include_dir = []
  include_ext = ["go", "tpl", "tmpl", "html", "yaml", "yml"]
  include_file = []
  kill_delay = "0s"
  log = "build-errors.log"
  poll = false
  poll_interval = 0
  rerun = false
  rerun_delay = 500
  send_interrupt = false
  stop_on_root = false

[color]
  app = ""
  build = "yellow"
  main = "magenta"
  runner = "green"
  watcher = "cyan"

[log]
  main_only = false
  time = false

[misc]
  clean_on_exit = false

[screen]
  clear_on_rebuild = false
  keep_scroll = true
`
}

func GitIgnore() string {
	return `# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with ` + "`go test -c`" + `
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Dependency directories (remove the comment below to include it)
# vendor/

# Go workspace file
go.work

# IDE files
.vscode/
.idea/
*.swp
*.swo
*~

# OS generated files
.DS_Store
.DS_Store?
._*
.Spotlight-V100
.Trashes
ehthumbs.db
Thumbs.db

# Application specific
tmp/
bin/
dist/
logs/
*.log
uploads/

# Environment files
.env
.env.local
.env.*.local

# Air
.air.toml.local

# Database
*.db
*.sqlite
*.sqlite3

# Coverage
coverage.out
coverage.html

# Build artifacts
build/
`
}

func Dockerfile() string {
	return `# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install dependencies
RUN apk add --no-cache git

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/configs ./configs

# Create uploads directory
RUN mkdir -p uploads

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Run the application
CMD ["./main"]
`
}

func DockerCompose(projectName string) string {
	return fmt.Sprintf(`version: '3.8'

services:
  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      - APP_ENV=production
      - DATABASE_HOST=postgres
      - DATABASE_USER=postgres
      - DATABASE_PASSWORD=password
      - DATABASE_NAME=%s_db
      - REDIS_HOST=redis
    depends_on:
      - postgres
      - redis
    volumes:
      - ./uploads:/root/uploads
    restart: unless-stopped

  postgres:
    image: postgres:15-alpine
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=password
      - POSTGRES_DB=%s_db
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
    restart: unless-stopped

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data
    restart: unless-stopped

  prometheus:
    image: prom/prometheus:latest
    ports:
      - "9090:9090"
    volumes:
      - ./deployments/prometheus.yml:/etc/prometheus/prometheus.yml
    restart: unless-stopped

  grafana:
    image: grafana/grafana:latest
    ports:
      - "3000:3000"
    environment:
      - GF_SECURITY_ADMIN_PASSWORD=admin
    volumes:
      - grafana_data:/var/lib/grafana
    restart: unless-stopped

volumes:
  postgres_data:
  redis_data:
  grafana_data:
`, projectName, projectName)
}