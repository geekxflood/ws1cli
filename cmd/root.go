// cmd/root.go

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var insecure bool // global variable to store the flag value
var prettyPrint bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ws1cli",
	Short: "CLI to interact with VMware Workspace ONE UEM API",
	Long:  `Command line interface to interact with VMware Workspace ONE UEM API.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
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
