package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd is the base/root command (i.e., `lazy`)
var rootCmd = &cobra.Command{
	Use:   "lazy",          // command name
	Short: "Lazy CLI tool", // shown in help output
	Run: func(cmd *cobra.Command, args []string) {
		// This runs when no subcommand is given
	},
}

// Execute is called from main.go to run the CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
