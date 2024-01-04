package yoyo

import (
	"fmt"
	"os"

	"github.com/nickpalenchar/yoyo/pkg/config"
	"github.com/spf13/cobra"
)

// NewYoyoCommand creates a new yoyo command
func NewDefaultYoyoCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "yoyo",
		Short: "Run the yoyo command",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO command logic
			Run()
		},
	}

	return cmd
}

func Run() {
	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}

	// Get yoyo configuration
	yoyoConfig, err := config.GetYoyoConfig()
	if err != nil {
		fmt.Println("Error getting yoyo configuration:", err)
		return
	}

	// Find matching commands for the current directory
	if commands, ok := yoyoConfig.Commands[cwd]; ok {
		fmt.Println("Available commands in", cwd+":")
		for key := range commands {
			fmt.Println("-", key)
		}
	} else {
		fmt.Println("No commands found for the current directory:", cwd)
	}
}
