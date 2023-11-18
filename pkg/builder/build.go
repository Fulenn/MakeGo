package builder

import (
	"fmt"
	"os"
	"path/filepath"
)

func Build(projectName string) {
	baseDir := filepath.Join("./", projectName)

	// Directories to create within the new project
	dirs := []string{"cmd", "pkg", "internal", "api", "configs"}

	// Loop through the directories and create them
	for _, dir := range dirs {
		dirPath := filepath.Join(baseDir, dir)
		err := os.MkdirAll(dirPath, 0755) // 0755 permission for directories
		if err != nil {
			fmt.Printf("Failed to create directory %s: %s\n", dirPath, err)
			return
		}
		fmt.Printf("Created directory %s\n", dirPath)
	}

	// Add any additional file creation or project setup logic here

	fmt.Println("Project created successfully:", projectName)
}
