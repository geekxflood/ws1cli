// cmd/root.go

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd is the base command for ws1cli, used when no subcommands are called.
var rootCmd = &cobra.Command{
	Use:   "ws1cli",
	Short: "CLI to interact with VMware Workspace ONE UEM API",
	Long:  `Command line interface to interact with VMware Workspace ONE UEM API.`,
}

// Execute adds all child commands to the root command.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Persistent flag available to all subcommands
	rootCmd.PersistentFlags().BoolVarP(&insecure, "insecure", "i", false, "Ignore TLS verification")
	rootCmd.PersistentFlags().BoolVarP(&prettyPrint, "pretty", "p", false, "Pretty-print JSON output")
}
