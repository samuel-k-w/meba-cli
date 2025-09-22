package generator

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func initGitRepo(projectPath, projectName string) error {
	// Initialize git repository
	cmd := exec.Command("git", "init")
	cmd.Dir = projectPath
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to initialize git repository: %w", err)
	}

	// Get git user info
	gitUser := getGitUser()
	if gitUser == "" {
		gitUser = "github.com/user"
	}

	// Update go.mod with proper module name
	moduleName := fmt.Sprintf("%s/%s", gitUser, projectName)
	if err := updateGoMod(projectPath, moduleName); err != nil {
		return fmt.Errorf("failed to update go.mod: %w", err)
	}
	
	// Update main.go with correct import path
	if err := updateMainGo(projectPath, moduleName); err != nil {
		return fmt.Errorf("failed to update main.go: %w", err)
	}

	// Add all files
	cmd = exec.Command("git", "add", ".")
	cmd.Dir = projectPath
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to add files to git: %w", err)
	}

	// Initial commit
	cmd = exec.Command("git", "commit", "-m", "Initial commit")
	cmd.Dir = projectPath
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to create initial commit: %w", err)
	}

	return nil
}

func getGitUser() string {
	// Try to get git remote origin URL
	cmd := exec.Command("git", "config", "--get", "remote.origin.url")
	if output, err := cmd.Output(); err == nil {
		url := strings.TrimSpace(string(output))
		if strings.Contains(url, "github.com") {
			// Extract username from GitHub URL
			if strings.HasPrefix(url, "git@github.com:") {
				parts := strings.Split(strings.TrimPrefix(url, "git@github.com:"), "/")
				if len(parts) > 0 {
					return "github.com/" + parts[0]
				}
			} else if strings.HasPrefix(url, "https://github.com/") {
				parts := strings.Split(strings.TrimPrefix(url, "https://github.com/"), "/")
				if len(parts) > 0 {
					return "github.com/" + parts[0]
				}
			}
		}
	}

	// Try to get git user name
	cmd = exec.Command("git", "config", "--get", "user.name")
	if output, err := cmd.Output(); err == nil {
		username := strings.TrimSpace(string(output))
		if username != "" {
			return "github.com/" + strings.ToLower(strings.ReplaceAll(username, " ", ""))
		}
	}

	return ""
}

func updateGoMod(projectPath, moduleName string) error {
	goModPath := fmt.Sprintf("%s/go.mod", projectPath)
	content, err := os.ReadFile(goModPath)
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")
	if len(lines) > 0 && strings.HasPrefix(lines[0], "module ") {
		lines[0] = fmt.Sprintf("module %s", moduleName)
	}

	updatedContent := strings.Join(lines, "\n")
	return os.WriteFile(goModPath, []byte(updatedContent), 0644)
}

func updateMainGo(projectPath, moduleName string) error {
	mainGoPath := fmt.Sprintf("%s/cmd/server/main.go", projectPath)
	content, err := os.ReadFile(mainGoPath)
	if err != nil {
		return err
	}

	projectName := strings.Split(moduleName, "/")[len(strings.Split(moduleName, "/"))-1]
	
	// Replace the import paths
	updatedContent := strings.ReplaceAll(string(content), 
		fmt.Sprintf("\"%s/internal\"", projectName),
		fmt.Sprintf("\"%s/internal\"", moduleName))
	
	// Fix docs import path
	updatedContent = strings.ReplaceAll(updatedContent,
		fmt.Sprintf("_ \"%s/docs\"", projectName),
		fmt.Sprintf("_ \"%s/docs\"", moduleName))

	return os.WriteFile(mainGoPath, []byte(updatedContent), 0644)
}

func runGoModTidy(projectPath string) error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = projectPath
	return cmd.Run()
}