package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Prefill terminal with echo hello",
	Long:  "Prefill the terminal cursor with 'echo \"hello\"'",
	Run: func(cmd *cobra.Command, args []string) {
		prefillTerminal(`echo "hello"`)
	},
}

func prefillTerminal(text string) {
	// Try iTerm2 first
	iTermScript := `tell application "iTerm" to activate
tell application "System Events"
	keystroke "` + text + `"
end tell`

	if err := exec.Command("osascript", "-e", iTermScript).Run(); err == nil {
		return
	}

	// Fall back to Terminal.app
	terminalScript := `tell application "Terminal" to activate
tell application "System Events"
	keystroke "` + text + `"
end tell`

	if err := exec.Command("osascript", "-e", terminalScript).Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: Could not prefill terminal. Make sure you're using Terminal.app or iTerm2.\n")
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
