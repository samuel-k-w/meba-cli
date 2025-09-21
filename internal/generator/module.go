package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/meba-cli/meba/internal/templates"
)

func GenerateModule(name string, dryRun, flat bool) error {
	var modulePath string
	
	if flat {
		modulePath = "."
	} else {
		modulePath = filepath.Join("internal", name)
	}

	if dryRun {
		fmt.Printf("Would create module at: %s\n", modulePath)
		return nil
	}

	// Create module directory
	if err := os.MkdirAll(modulePath, 0755); err != nil {
		return fmt.Errorf("failed to create module directory: %w", err)
	}

	// Generate module files
	files := map[string]string{
		"module.go":     templates.ModuleGo(name),
		"handlers.go":   templates.ModuleHandlersGo(name),
		"service.go":    templates.ModuleServiceGo(name),
		"repository.go": templates.ModuleRepositoryGo(name),
		"entity.go":     templates.ModuleEntityGo(name),
		"dto.go":        templates.ModuleDtoGo(name),
	}

	for fileName, content := range files {
		filePath := filepath.Join(modulePath, fileName)
		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to write file %s: %w", fileName, err)
		}
	}

	// Update app.go to register the module
	if !flat {
		if err := updateAppModule(name); err != nil {
			fmt.Printf("Warning: Could not auto-register module in app.go: %v\n", err)
		}
	}

	return nil
}

func GenerateHandler(name string, dryRun, flat, noSpec bool) error {
	modulePath, exists := findModulePath(name)
	if !exists && !flat {
		modulePath = filepath.Join("internal", name)
		os.MkdirAll(modulePath, 0755)
	} else if flat {
		modulePath = "."
	}

	if dryRun {
		fmt.Printf("Would create handler at: %s/handlers.go\n", modulePath)
		return nil
	}

	content := templates.ModuleHandlersGo(name)
	filePath := filepath.Join(modulePath, "handlers.go")
	
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write handler file: %w", err)
	}

	// Update module.go to include handler
	if err := updateModuleFile(modulePath, name, "handler"); err != nil {
		fmt.Printf("Warning: Could not auto-register handler: %v\n", err)
	}

	return nil
}

func GenerateService(name string, dryRun, flat, noSpec bool) error {
	modulePath, exists := findModulePath(name)
	if !exists && !flat {
		modulePath = filepath.Join("internal", name)
		os.MkdirAll(modulePath, 0755)
	} else if flat {
		modulePath = "."
	}

	if dryRun {
		fmt.Printf("Would create service at: %s/service.go\n", modulePath)
		return nil
	}

	content := templates.ModuleServiceGo(name)
	filePath := filepath.Join(modulePath, "service.go")
	
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write service file: %w", err)
	}

	// Update module.go to include service
	if err := updateModuleFile(modulePath, name, "service"); err != nil {
		fmt.Printf("Warning: Could not auto-register service: %v\n", err)
	}

	return nil
}

func GenerateRepository(name string, dryRun, flat, noSpec bool) error {
	modulePath, exists := findModulePath(name)
	if !exists && !flat {
		modulePath = filepath.Join("internal", name)
		os.MkdirAll(modulePath, 0755)
	} else if flat {
		modulePath = "."
	}

	if dryRun {
		fmt.Printf("Would create repository at: %s/repository.go\n", modulePath)
		return nil
	}

	content := templates.ModuleRepositoryGo(name)
	filePath := filepath.Join(modulePath, "repository.go")
	
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write repository file: %w", err)
	}

	// Update module.go to include repository
	if err := updateModuleFile(modulePath, name, "repository"); err != nil {
		fmt.Printf("Warning: Could not auto-register repository: %v\n", err)
	}

	return nil
}

func GenerateResource(name string, dryRun bool) error {
	if dryRun {
		fmt.Printf("Would create complete CRUD resource for: %s\n", name)
		return nil
	}

	// Generate module first
	if err := GenerateModule(name, false, false); err != nil {
		return err
	}

	// Generate all components
	if err := GenerateHandler(name, false, false, false); err != nil {
		return err
	}
	
	if err := GenerateService(name, false, false, false); err != nil {
		return err
	}
	
	if err := GenerateRepository(name, false, false, false); err != nil {
		return err
	}

	return nil
}

func GenerateMiddleware(name string, dryRun, flat bool) error {
	var filePath string
	
	if flat {
		filePath = fmt.Sprintf("%s_middleware.go", name)
	} else {
		filePath = filepath.Join("pkg", "middleware", fmt.Sprintf("%s.go", name))
		os.MkdirAll(filepath.Dir(filePath), 0755)
	}

	if dryRun {
		fmt.Printf("Would create middleware at: %s\n", filePath)
		return nil
	}

	content := templates.MiddlewareGo(name)
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write middleware file: %w", err)
	}

	return nil
}

func GenerateGuard(name string, dryRun, flat bool) error {
	var filePath string
	
	if flat {
		filePath = fmt.Sprintf("%s_guard.go", name)
	} else {
		filePath = filepath.Join("pkg", "middleware", fmt.Sprintf("%s_guard.go", name))
		os.MkdirAll(filepath.Dir(filePath), 0755)
	}

	if dryRun {
		fmt.Printf("Would create guard at: %s\n", filePath)
		return nil
	}

	content := templates.GuardGo(name)
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write guard file: %w", err)
	}

	return nil
}

func updateModuleFile(modulePath, moduleName, componentType string) error {
	moduleFilePath := filepath.Join(modulePath, "module.go")
	
	// Check if module.go exists
	if _, err := os.Stat(moduleFilePath); os.IsNotExist(err) {
		// Create module.go if it doesn't exist
		content := templates.ModuleGo(moduleName)
		return os.WriteFile(moduleFilePath, []byte(content), 0644)
	}

	content, err := os.ReadFile(moduleFilePath)
	if err != nil {
		return err
	}

	updatedContent := string(content)
	
	// Add component to wire.NewSet
	componentName := fmt.Sprintf("New%s%s", strings.Title(moduleName), strings.Title(componentType))
	if !strings.Contains(updatedContent, componentName) {
		// Find wire.NewSet and add component
		wireIndex := strings.Index(updatedContent, "wire.NewSet(")
		if wireIndex != -1 {
			endIndex := strings.Index(updatedContent[wireIndex:], ")")
			if endIndex != -1 {
				insertPos := wireIndex + endIndex
				addition := fmt.Sprintf("\n\t%s,", componentName)
				updatedContent = updatedContent[:insertPos] + addition + updatedContent[insertPos:]
			}
		}
	}

	return os.WriteFile(moduleFilePath, []byte(updatedContent), 0644)
}