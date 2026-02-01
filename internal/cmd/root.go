package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ct",
	Short: "Terminal command prefiller",
	Long:  "ct prefills your terminal with commonly used commands",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ct v" + Version)
		fmt.Println("Run 'ct --help' for usage information")
	},
}

var Version = "0.1.0"

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("help", "h", false, "help for ct")
}
