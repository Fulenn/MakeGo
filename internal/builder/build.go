package builder

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func Build(projectName string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Do you want to create a full project? (yes/no): ")
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(response) // Remove any trailing newline character

	// Define main and full directories
	mainDirs := []string{"cmd", "internal", "pkg", "vendor"}
	fullDirs := []string{"api", "web", "configs", "init", "scripts", "build", "deployments", "test"}

	// Select directories based on user response
	var dirs []string
	if strings.ToLower(response) == "yes" {
		dirs = append(mainDirs, fullDirs...)
	} else {
		dirs = mainDirs
	}

	baseDir := filepath.Join("./", projectName)

	if err := createDirectories(baseDir, dirs); err != nil {
		fmt.Println(err)
		return
	}

	// Ensure the cmd/projectName directory exists before creating main.go
	cmdDirPath := filepath.Join(baseDir, "cmd")
	if err := os.MkdirAll(cmdDirPath, 0755); err != nil {
		fmt.Printf("Failed to create directory %s: %s\n", cmdDirPath, err)
		return
	}

	mainGoContent, err := loadMainGoTemplate() // Load from file
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create main.go file
	mainGoPath := filepath.Join(cmdDirPath, "main.go")
	if err := ioutil.WriteFile(mainGoPath, []byte(mainGoContent), 0644); err != nil {
		fmt.Printf("Failed to create main.go file: %s\n", err)
		return
	}
	fmt.Printf("Created main.go in %s\n", cmdDirPath)

	fmt.Println("Project created successfully:", projectName)
}

func createDirectories(baseDir string, dirs []string) error {
	for _, dir := range dirs {
		dirPath := filepath.Join(baseDir, dir)
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %s", dirPath, err)
		}
		fmt.Printf("Created directory %s\n", dirPath)
	}
	return nil
}

func loadMainGoTemplate() (string, error) {
	templatePath := "./internal/builder/templates/main.go.tpl"
	content, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to load main.go template: %s", err)
	}
	return string(content), nil
}
