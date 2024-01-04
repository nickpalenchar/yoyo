package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
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

// GetYoyoConfig reads the config from $HOME/.yoyo/yoyo.json and returns it.
// If the file doesn't exist, it creates and saves a default config to $HOME/.yoyo/yoyo.json, and returns that default.
func GetYoyoConfig() (*YoyoConfig, error) {
	// Get the home directory
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("error getting user's home directory: %w", err)
	}

	// Build the path to yoyo.json
	configPath := filepath.Join(home, ".yoyo", "yoyo.json")

	// Read yoyo.json file or create a default one if it doesn't exist
	configFile, err := os.ReadFile(configPath)
	var yoyoConfig *YoyoConfig

	if os.IsNotExist(err) {
		// Create default config
		yoyoConfig = DefaultYoyoConfig()
		if err := saveYoyoConfig(configPath, yoyoConfig); err != nil {
			return nil, fmt.Errorf("error saving default yoyo.json: %w", err)
		}
	} else if err != nil {
		return nil, fmt.Errorf("error reading yoyo.json: %w", err)
	} else {
		// Parse existing config
		if err := json.Unmarshal(configFile, &yoyoConfig); err != nil {
			return nil, fmt.Errorf("error parsing yoyo.json: %w", err)
		}
	}

	return yoyoConfig, nil
}

// DefaultYoyoConfig returns a default YoyoConfig with an empty commands map
func DefaultYoyoConfig() *YoyoConfig {
	return &YoyoConfig{
		Commands: make(DirectoryCommands),
	}
}

// saveYoyoConfig saves the YoyoConfig to yoyo.json
func saveYoyoConfig(configPath string, config *YoyoConfig) error {
	// Create intermediate directories if they don't exist
	dir := filepath.Dir(configPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("error creating intermediate directories: %w", err)
	}

	// Marshal the YoyoConfig
	file, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling yoyo.json: %w", err)
	}

	// Write the file
	if err := os.WriteFile(configPath, file, 0644); err != nil {
		return fmt.Errorf("error writing yoyo.json: %w", err)
	}

	return nil
}
