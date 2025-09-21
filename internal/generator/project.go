package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/meba-cli/meba/internal/templates"
)

func CreateProject(name, targetDir string, skipGit bool) error {
	// Create project directory
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Create directory structure
	dirs := []string{
		"cmd/server",
		"internal",
		"pkg/middleware",
		"pkg/validator",
		"configs",
		"scripts",
		"deployments",
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(filepath.Join(targetDir, dir), 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	// Get module name (will be updated by git init)
	moduleName := name
	
	// Generate files
	files := map[string]string{
		"go.mod":                    templates.GoMod(moduleName),
		"README.md":                 templates.ReadmeMd(name),
		"cmd/server/main.go":        templates.MainGo(moduleName),
		"internal/app.go":           templates.AppGo(),
		"internal/handlers.go":      templates.HandlersGo(),
		"internal/service.go":       templates.ServiceGo(),
		"internal/entity.go":        templates.EntityGo(),
		"internal/dto.go":           templates.DtoGo(),
		"internal/repository.go":    templates.RepositoryGo(),
		"internal/wire.go":         templates.WireGo(),
		"pkg/middleware/logger.go":  templates.LoggerMiddleware(),
		"pkg/middleware/recover.go": templates.RecoverMiddleware(),
		"pkg/middleware/auth.go":    templates.AuthMiddleware(),
		"pkg/validator/validator.go": templates.ValidatorGo(),
		"configs/config.yaml":       templates.ConfigYaml(),
		"configs/config.go":         templates.ConfigGo(),
		".air.toml":                 templates.AirToml(),
		".gitignore":                templates.GitIgnore(),
		"Dockerfile":                templates.Dockerfile(),
		"docker-compose.yml":        templates.DockerCompose(name),
	}

	for filePath, content := range files {
		fullPath := filepath.Join(targetDir, filePath)
		if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
			return fmt.Errorf("failed to create directory for %s: %w", filePath, err)
		}
		if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to write file %s: %w", filePath, err)
		}
	}

	// Initialize git repository
	if !skipGit {
		if err := initGitRepo(targetDir, name); err != nil {
			fmt.Printf("Warning: Could not initialize git repository: %v\n", err)
		}
	}

	// Run go mod tidy
	if err := runGoModTidy(targetDir); err != nil {
		fmt.Printf("Warning: Could not run go mod tidy: %v\n", err)
	}

	return nil
}

func findModulePath(name string) (string, bool) {
	// Check if we're in a module directory
	currentDir, _ := os.Getwd()
	
	// Look for existing module directories
	entries, err := os.ReadDir(filepath.Join(currentDir, "internal"))
	if err != nil {
		return "", false
	}

	for _, entry := range entries {
		if entry.IsDir() && strings.ToLower(entry.Name()) == strings.ToLower(name) {
			return filepath.Join("internal", entry.Name()), true
		}
	}

	return "", false
}

func updateAppModule(moduleName string) error {
	appPath := "internal/app.go"
	content, err := os.ReadFile(appPath)
	if err != nil {
		return err
	}

	updatedContent := string(content)
	
	// Add import
	importLine := fmt.Sprintf("\"%s/internal/%s\"", getCurrentModuleName(), moduleName)
	if !strings.Contains(updatedContent, importLine) {
		importIndex := strings.Index(updatedContent, "import (")
		if importIndex != -1 {
			endIndex := strings.Index(updatedContent[importIndex:], ")")
			if endIndex != -1 {
				insertPos := importIndex + endIndex
				updatedContent = updatedContent[:insertPos] + "\n\t" + importLine + updatedContent[insertPos:]
			}
		}
	}

	// Add module registration
	regLine := fmt.Sprintf("%s.Module,", moduleName)
	if !strings.Contains(updatedContent, regLine) {
		wireIndex := strings.Index(updatedContent, "AppSet = wire.NewSet(")
		if wireIndex != -1 {
			endIndex := strings.Index(updatedContent[wireIndex:], ")")
			if endIndex != -1 {
				insertPos := wireIndex + endIndex
				updatedContent = updatedContent[:insertPos] + "\n\t" + regLine + updatedContent[insertPos:]
			}
		}
	}

	return os.WriteFile(appPath, []byte(updatedContent), 0644)
}

func getCurrentModuleName() string {
	content, err := os.ReadFile("go.mod")
	if err != nil {
		return "myapp"
	}
	
	lines := strings.Split(string(content), "\n")
	if len(lines) > 0 && strings.HasPrefix(lines[0], "module ") {
		return strings.TrimSpace(strings.TrimPrefix(lines[0], "module "))
	}
	
	return "myapp"
}