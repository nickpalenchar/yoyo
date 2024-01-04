package yoyo

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// Command structure represents a single command entry
type Command struct {
	Cmd         string `json:"cmd"`
	Description string `json:"description,omitempty"`
}

// CommandsMap is a map of command names to Command structures
type CommandsMap map[string]Command

// DirectoryCommands is a map of directory paths to CommandsMap
type DirectoryCommands map[string]CommandsMap

// YoyoConfig represents the overall configuration structure
type YoyoConfig struct {
	Commands DirectoryCommands `json:"commands"`
}

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

// DefaultYoyoConfig returns a default YoyoConfig with an empty commands map
func DefaultYoyoConfig() *YoyoConfig {
	return &YoyoConfig{
		Commands: make(DirectoryCommands),
	}
}

func Run() {
	// Get the home directory
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user's home directory:", err)
		return
	}

	// Build the path to yoyo.json
	configPath := filepath.Join(home, ".yoyo", "yoyo.json")

	// Read yoyo.json file or create a default one if it doesn't exist
	configFile, err := os.ReadFile(configPath)
	var yoyoConfig *YoyoConfig

	if os.IsNotExist(err) {
		// Create default config
		yoyoConfig = DefaultYoyoConfig()
	} else if err != nil {
		fmt.Println("Error reading yoyo.json:", err)
		return
	} else {
		// Parse existing config
		if err := json.Unmarshal(configFile, &yoyoConfig); err != nil {
			fmt.Println("Error parsing yoyo.json:", err)
			return
		}
	}

	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
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
