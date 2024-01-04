package add

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// NewAddCommand creates a new add command
func NewAddCommand() *cobra.Command {
	var (
		name        string
		cmd         string
		description string
	)

	var cmdAdd = &cobra.Command{
		Use:   "add",
		Short: "Echo back user-inputted values",
		Run: func(_ *cobra.Command, _ []string) {
			// Prompt for values if not provided as flags
			if name == "" {
				name = getUserInput("Enter name:")
			}
			if cmd == "" {
				cmd = getUserInput("Enter cmd:")
			}

			// Echo back user-inputted values
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
