package cmd

import (
	"fmt"
	"github.com/Fulenn/MakeGo/internal/builder"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   `makego`,
	Short: `MakeGo is a CLI tool for generating Go projects`,
	Long:  `MakeGo is a CLI tool for generating Go projects based on a predefined layout.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			builder.Build(args[0]) // Passes the first argument as the project name
		} else {
			fmt.Println("Please provide a project name")
			// Optionally, you can return an error or exit the command if no argument is provided
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
