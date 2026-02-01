package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Copy text to clipboard",
	Long:  "Copy provided text to clipboard and display it",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Fprintf(os.Stderr, "Error: Please provide text to copy\n")
			os.Exit(1)
		}
		text := strings.Join(args, " ")
		copyToClipboard(text)
		// fmt.Printf(">>[copied to clipboard>> %s\n", text)
		fmt.Printf("\033[32m✅ Copied to clipboard:\033[0m %s", text)
	},
}

func copyToClipboard(text string) {
	cmd := exec.Command("pbcopy")
	cmd.Stdin = strings.NewReader(text)
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: Could not copy to clipboard\n")
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
