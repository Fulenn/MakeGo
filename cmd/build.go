package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Builds a new project",
	Long: `The build command creates a new project directory 
based on a predefined layout. Usage:

ProjectBuilder build [project name]`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			build(args[0]) // Passes the first argument as the project name
		} else {
			fmt.Println("Please provide a project name")
			// Optionally, you can return an error or exit the command if no argument is provided
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

func build(projectName string) {
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
