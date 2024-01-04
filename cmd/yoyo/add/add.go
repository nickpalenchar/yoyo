package add

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const yoyoConfigPath = "$HOME/.yoyo/yoyo.json"

// NewAddCommand creates a new add command
func NewAddCommand() *cobra.Command {
	var (
		name        string
		cmd         string
		description string
	)

	var cmdAdd = &cobra.Command{
		Use:   "add",
		Short: "Add a new command to $HOME/.yoyo/yoyo.json",
		Run: func(_ *cobra.Command, _ []string) {
			// Prompt for values if not provided as flags
			if name == "" {
				name = getUserInput("Enter name:")
			}
			if cmd == "" {
				cmd = getUserInput("Enter cmd:")
			}
			if description == "" {
				description = getUserInput("Enter description:")
			}

			// Read existing yoyo.json content
			configPath := os.ExpandEnv(yoyoConfigPath)
			config, err := readYoyoConfig(configPath)
			if err != nil {
				fmt.Println("Error reading yoyo.json:", err)
				return
			}

			// Add the new command to the config
			addCommandToConfig(config, name, cmd, description)

			// Save the updated config back to yoyo.json
			err = saveYoyoConfig(configPath, config)
			if err != nil {
				fmt.Println("Error saving yoyo.json:", err)
				return
			}

			// Echo back user-inputted values
			fmt.Println("Command added successfully:")
			fmt.Println("Name:", name)
			fmt.Println("Command:", cmd)
			fmt.Println("Description:", description)
		},
	}

	// Add flags to the command
	cmdAdd.Flags().StringVarP(&name, "name", "n", "", "Name of the command")
	cmdAdd.Flags().StringVarP(&cmd, "cmd", "c", "", "Command to execute")
	cmdAdd.Flags().StringVarP(&description, "description", "d", "", "Description of the command")

	return cmdAdd
}

// getUserInput prompts the user for input and returns the entered value
func getUserInput(prompt string) string {
	fmt.Print(prompt + " ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

// readYoyoConfig reads the existing yoyo.json content
func readYoyoConfig(configPath string) (map[string]interface{}, error) {
	file, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config map[string]interface{}
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// addCommandToConfig adds the new command to the yoyo.json config
func addCommandToConfig(config map[string]interface{}, name, cmd, description string) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	commands, ok := config["commands"].(map[string]interface{})
	if !ok {
		commands = make(map[string]interface{})
		config["commands"] = commands
	}

	commands[dir] = map[string]interface{}{
		"name":        name,
		"cmd":         cmd,
		"description": description,
	}
}

// saveYoyoConfig saves the updated yoyo.json content
func saveYoyoConfig(configPath string, config map[string]interface{}) error {
	file, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, file, 0644)
	if err != nil {
		return err
	}

	return nil
}
