package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var rawOutput bool

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
			os.Exit(0)
		}

		// Trim whitespace
		command = strings.TrimSpace(command)

		// If raw output is requested, just print the command and exit
		if rawOutput {
			fmt.Print(command)
			return
		}

		// Copy to clipboard
		pbcopyCmd := exec.Command("pbcopy")
		pbcopyCmd.Stdin = strings.NewReader(command)
		if err := pbcopyCmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: Could not copy to clipboard\n")
			os.Exit(0)
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

	// Handle --raw or -r flag at the beginning
	if len(args) > 0 && (args[0] == "--raw" || args[0] == "-r") {
		// Check if there are more arguments after the flag
		if len(args) > 1 {
			// Treat the rest as a query
			userQuery := strings.Join(args[1:], " ")
			processQueryWithRaw(userQuery)
			return
		}
		// If no query provided with --raw, fall through to normal execution for error
	}

	// Check if it's a known subcommand
	if len(args) > 0 {
		if args[0] != "set-key" && args[0] != "qq-install" && args[0] != "qq-uninstall" && args[0] != "help" && args[0] != "-h" && args[0] != "--help" {
			// Treat as a query
			userQuery := strings.Join(args, " ")
			processQuery(userQuery)
			return
		}
	}

	// Execute as normal cobra command
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(0)
	}
}

func processQuery(userQuery string) {
	// Call GROQ API to get the command
	command, err := callGroqAPI(userQuery)
	if err != nil {
		// Print error directly (to stdout for proper ANSI color rendering)
		fmt.Print(err)
		os.Exit(0)
	}

	// Trim whitespace
	command = strings.TrimSpace(command)

	// Copy to clipboard
	pbcopyCmd := exec.Command("pbcopy")
	pbcopyCmd.Stdin = strings.NewReader(command)
	if err := pbcopyCmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: Could not copy to clipboard\n")
		os.Exit(0)
	}

	// Print the command
	fmt.Printf("\033[32m✅ The command has been copied to your clipboard and is ready to paste:\033[0m\n")
	fmt.Printf("\033[1m%s\033[0m\n", command)
}

func processQueryWithRaw(userQuery string) {
	// Call GROQ API to get the command
	command, err := callGroqAPI(userQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(0)
	}

	// Trim whitespace and output only the raw command
	command = strings.TrimSpace(command)
	fmt.Print(command)
}

func init() {
	rootCmd.Flags().BoolP("help", "h", false, "help for qx")
	rootCmd.Flags().BoolVarP(&rawOutput, "raw", "r", false, "output only the command without formatting")
}
