package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
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

func build(s string) {
	fmt.Println("Building project", s)
}
