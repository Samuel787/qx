package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qx",
	Short: "Terminal command prefiller",
	Long:  "qx prefills your terminal with commonly used commands",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("qx v" + Version)
			fmt.Println("Run 'qx --help' for usage information")
			return
		}

		// Join all arguments into a query
		userQuery := strings.Join(args, " ")

		// Call GROQ API to get the command
		command, err := callGroqAPI(userQuery)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		// Trim whitespace
		command = strings.TrimSpace(command)

		// Copy to clipboard
		pbcopyCmd := exec.Command("pbcopy")
		pbcopyCmd.Stdin = strings.NewReader(command)
		if err := pbcopyCmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: Could not copy to clipboard\n")
			os.Exit(1)
		}

		// Print the command
		fmt.Printf("\033[32m✅ The command has been copied to your clipboard and is ready to paste:\033[0m\n")
		fmt.Printf("\033[1m%s\033[0m\n", command)
	},
	SilenceUsage:  true,
	SilenceErrors: true,
}

var Version = "0.1.0"

func Execute() {
	// Try to execute as normal first
	args := os.Args[1:]

	// Check if it's a known subcommand
	if len(args) > 0 {
		if args[0] != "set-key" && args[0] != "help" && args[0] != "-h" && args[0] != "--help" {
			// Treat as a query
			userQuery := strings.Join(args, " ")
			processQuery(userQuery)
			return
		}
	}

	// Execute as normal cobra command
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func processQuery(userQuery string) {
	// Call GROQ API to get the command
	command, err := callGroqAPI(userQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Trim whitespace
	command = strings.TrimSpace(command)

	// Copy to clipboard
	pbcopyCmd := exec.Command("pbcopy")
	pbcopyCmd.Stdin = strings.NewReader(command)
	if err := pbcopyCmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: Could not copy to clipboard\n")
		os.Exit(1)
	}

	// Print the command
	fmt.Printf("\033[32m✅ The command has been copied to your clipboard and is ready to paste:\033[0m\n")
	fmt.Printf("\033[1m%s\033[0m\n", command)
}

func init() {
	rootCmd.Flags().BoolP("help", "h", false, "help for qx")
}
