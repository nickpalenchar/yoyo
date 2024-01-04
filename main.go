package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nickpalenchar/yoyo/cmd/yoyo"
	"github.com/nickpalenchar/yoyo/cmd/yoyo/add"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "yoyo",
		Short: "yoyo - run commands",
		Run: func(cmd *cobra.Command, args []string) {
			// Run the yoyo command by default when no arguments are provided
			yoyo.Run()
		},
	}

	// Attach yoyo command to root command
	rootCmd.AddCommand(yoyo.NewDefaultYoyoCommand())
	rootCmd.AddCommand(add.NewAddCommand())

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		// Handle errors
		panic(err)
	}
}

func main_old() {
	fmt.Println("Hello, Go World!")

	// Your application logic goes here

	// Example: Read an environment variable
	envVarValue := os.Getenv("MY_ENV_VAR")
	fmt.Printf("Value of MY_ENV_VAR: %s\n", envVarValue)

	yoyo.Run()

	// Example: Use a logger
	log.Println("This is a log message")

	// Example: Call a function from another package
	// mypackage.MyFunction()

	// Exit the application
	os.Exit(0)
}
