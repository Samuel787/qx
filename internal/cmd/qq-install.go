package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var qqInstallCmd = &cobra.Command{
	Use:   "qq-install",
	Short: "Install qq() function to your shell",
	Long:  "Install the qq() shell function that allows you to use 'qq <query>' to get commands from Groq AI",
	Run: func(cmd *cobra.Command, args []string) {
		installQQFunction()
	},
}

func installQQFunction() {
	// Determine which shell config file to use
	shell := os.Getenv("SHELL")
	var rcFile string
	var sourceCmd string
	var shellType string

	if strings.Contains(shell, "zsh") {
		rcFile = filepath.Join(os.Getenv("HOME"), ".zshrc")
		sourceCmd = "source ~/.zshrc"
		shellType = "zsh"
	} else if strings.Contains(shell, "bash") {
		rcFile = filepath.Join(os.Getenv("HOME"), ".bashrc")
		sourceCmd = "source ~/.bashrc"
		shellType = "bash"
	} else {
		fmt.Fprintf(os.Stderr, "Error: Unsupported shell. Please use bash or zsh\n")
		os.Exit(1)
	}

	// Read existing content
	content, err := os.ReadFile(rcFile)
	if err != nil && !os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: Could not read %s\n", rcFile)
		os.Exit(1)
	}

	fileContent := string(content)

	// Check if qq function already exists
	if strings.Contains(fileContent, "qq()") || strings.Contains(fileContent, "function qq") {
		fmt.Printf("\033[33m⚠️  The qq() function already exists in your %s.\033[0m\n", filepath.Base(rcFile))
		fmt.Printf("If you want to reinstall, please manually remove the qq() function first.\n")
		return
	}

	// Create the qq() shell function
	var qqFunction string
	if shellType == "zsh" {
		qqFunction = `
# >>> qx tool >>>
# qq() function - Query Groq AI for shell commands
qq() {
	if [[ -z "$1" ]]; then
		echo "Usage: qq <query>"
		return 1
	fi

	# Call qx with the query and get raw output
	local result=$(qx --raw "$@" 2>&1)

	# Check if the command was successful
	if [[ $? -eq 0 ]]; then
		# Use zsh's print -z to prefill the command line
		print -z "$result"
	else
		echo "Error: $result"
		return 1
	fi
}
# <<< qx tool <<<
`
	} else {
		// Bash version
		qqFunction = `
# >>> qx tool >>>
# qq() function - Query Groq AI for shell commands
qq() {
	if [[ -z "$1" ]]; then
		echo "Usage: qq <query>"
		return 1
	fi

	# Call qx with the query and get raw output
	local result=$(qx --raw "$@" 2>&1)

	# Check if the command was successful
	if [[ $? -eq 0 ]]; then
		# Use printf to write the command to stdin for readline
		# This simulates pasting the command
		printf "%s" "$result" | xclip -selection clipboard 2>/dev/null || printf "%s" "$result" | pbcopy 2>/dev/null

		# Add to history for easy access with up arrow
		history -s "$result"

		# Display the command with instructions
		echo "Command ready (copied to clipboard): $result"
		echo "Press Ctrl+V (or Cmd+V on Mac) to paste, or press Up arrow to recall from history"
	else
		echo "Error: $result"
		return 1
	fi
}
# <<< qx tool <<<
`
	}

	// Append the function to the rc file
	if fileContent != "" && !strings.HasSuffix(fileContent, "\n") {
		fileContent += "\n"
	}
	fileContent += qqFunction + "\n"

	// Write back to file
	if err := os.WriteFile(rcFile, []byte(fileContent), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error: Could not write to %s\n", rcFile)
		os.Exit(1)
	}

	// Copy source command to clipboard
	pbcopyCmd := exec.Command("pbcopy")
	pbcopyCmd.Stdin = strings.NewReader(sourceCmd)
	if err := pbcopyCmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: Could not copy to clipboard\n")
		os.Exit(1)
	}

	// Inform user
	fmt.Printf("\033[32m✅ The qq() function has been installed!\033[0m\n\n")
	fmt.Printf("To start using it, run:\n")
	fmt.Printf("\033[1m%s\033[0m\n\n", sourceCmd)
	fmt.Printf("Or you can:\n")
	fmt.Printf("  1. Close this terminal and open a new one\n")
	fmt.Printf("  2. Start a new shell with: exec $SHELL\n")
	fmt.Printf("  3. Manually run: %s\n\n", sourceCmd)
	fmt.Printf("Then you can use: \033[1mqq <your query>\033[0m\n")
	fmt.Printf("Example: \033[1mqq list files in current directory\033[0m\n")
	fmt.Printf("\n💡 The source command has been copied to your clipboard!\n")
}

func init() {
	rootCmd.AddCommand(qqInstallCmd)
}
