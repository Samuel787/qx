package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var setApiKeyCmd = &cobra.Command{
	Use:   "set-key [token]",
	Short: "Set QX_GROQ_KEY environment variable",
	Long:  "Set the QX_GROQ_KEY environment variable in your shell configuration file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		token := args[0]
		setApiKey(token)
	},
}

func setApiKey(token string) {
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
		os.Exit(1)
	}

	// Read existing content
	content, err := os.ReadFile(rcFile)
	if err != nil && !os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: Could not read %s\n", rcFile)
		os.Exit(1)
	}

	// Check if QX_GROQ_KEY already exists
	fileContent := string(content)
	exportLine := fmt.Sprintf("export QX_GROQ_KEY=\"%s\"", token)

	if strings.Contains(fileContent, "QX_GROQ_KEY") {
		// Replace existing key
		lines := strings.Split(fileContent, "\n")
		var newLines []string
		for _, line := range lines {
			if strings.Contains(line, "QX_GROQ_KEY") {
				newLines = append(newLines, exportLine)
			} else {
				newLines = append(newLines, line)
			}
		}
		fileContent = strings.Join(newLines, "\n")
	} else {
		// Append new key
		if fileContent != "" && !strings.HasSuffix(fileContent, "\n") {
			fileContent += "\n"
		}
		fileContent += exportLine + "\n"
	}

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
	fmt.Printf("\033[32m✅ API key has been set. Now, run the following to activate the environment variable:\n\033[0m")
	// fmt.Printf("")
	fmt.Printf("\033[1m%s\033[0m\n", sourceCmd)
	// fmt.Printf("\033[36m💡 The source command has been copied to your clipboard, so just paste it!\033[0m\n")
}

func init() {
	rootCmd.AddCommand(setApiKeyCmd)
}
