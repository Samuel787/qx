package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var qqUninstallCmd = &cobra.Command{
	Use:   "qq-uninstall",
	Short: "Uninstall qq() function from your shell",
	Long:  "Remove the qq() shell function from your shell configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		uninstallQQFunction()
	},
}

func uninstallQQFunction() {
	// Determine which shell config file to use
	shell := os.Getenv("SHELL")
	var rcFile string
	var sourceCmd string

	if strings.Contains(shell, "zsh") {
		rcFile = filepath.Join(os.Getenv("HOME"), ".zshrc")
		sourceCmd = "source ~/.zshrc"
	} else if strings.Contains(shell, "bash") {
		rcFile = filepath.Join(os.Getenv("HOME"), ".bashrc")
		sourceCmd = "source ~/.bashrc"
	} else {
		fmt.Fprintf(os.Stderr, "Error: Unsupported shell. Please use bash or zsh\n")
		os.Exit(0)
	}

	// Read existing content
	content, err := os.ReadFile(rcFile)
	if err != nil && !os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: Could not read %s\n", rcFile)
		os.Exit(0)
	}

	fileContent := string(content)

	// Check if qq function exists
	if !strings.Contains(fileContent, "# >>> qx tool >>>") {
		fmt.Printf("\033[33m⚠️  The qq() function is not installed in your %s.\033[0m\n", filepath.Base(rcFile))
		return
	}

	// Remove the qq() function block between the markers
	fileContent = removeQQFunction(fileContent)

	// Write back to file
	if err := os.WriteFile(rcFile, []byte(fileContent), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error: Could not write to %s\n", rcFile)
		os.Exit(0)
	}

	// Copy source command to clipboard
	pbcopyCmd := exec.Command("pbcopy")
	pbcopyCmd.Stdin = strings.NewReader(sourceCmd)
	if err := pbcopyCmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: Could not copy to clipboard\n")
		os.Exit(0)
	}

	// Inform user
	fmt.Printf("\033[32m✅ The qq() function has been removed from your %s!\033[0m\n\n", filepath.Base(rcFile))
	fmt.Printf("To apply the changes, run:\n")
	fmt.Printf("\033[1m%s\033[0m\n\n", sourceCmd)
	fmt.Printf("Or you can:\n")
	fmt.Printf("  1. Close this terminal and open a new one\n")
	fmt.Printf("  2. Start a new shell with: exec $SHELL\n")
	fmt.Printf("  3. Manually run: %s\n\n", sourceCmd)
	fmt.Printf("💡 The source command has been copied to your clipboard!\n")
}

func removeQQFunction(content string) string {
	lines := strings.Split(content, "\n")
	var result []string
	inQXBlock := false

	for _, line := range lines {
		// Check if we're entering the qx tool block
		if strings.Contains(line, "# >>> qx tool >>>") {
			inQXBlock = true
			continue
		}

		// Check if we're exiting the qx tool block
		if strings.Contains(line, "# <<< qx tool <<<") {
			inQXBlock = false
			continue
		}

		// Only add lines that are not in the qx tool block
		if !inQXBlock {
			result = append(result, line)
		}
	}

	// Join lines back together
	finalContent := strings.Join(result, "\n")

	// Clean up excessive blank lines
	// Remove multiple consecutive blank lines, but keep single blank lines
	for strings.Contains(finalContent, "\n\n\n") {
		finalContent = strings.ReplaceAll(finalContent, "\n\n\n", "\n\n")
	}

	// Trim trailing whitespace but keep the structure
	finalContent = strings.TrimRight(finalContent, " \t\n")
	if finalContent != "" {
		finalContent += "\n"
	}

	return finalContent
}

func init() {
	rootCmd.AddCommand(qqUninstallCmd)
}
