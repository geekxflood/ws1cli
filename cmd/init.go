// cmd/init.go

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the 'init' command for initializing the CLI.
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the CLI",
	Long:  `This command sets up the CLI by guiding the user through the creation of a configuration file. It will check for an existing config and prompt to create one if not found.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Directly create a new config file if -f flag is used
		if forceRecreate {
			fmt.Println("Recreating configuration file...")
			if err := createConfigFile(); err != nil {
				fmt.Println("Error creating config:", err)
				os.Exit(1)
			}
			fmt.Println("Configuration file recreated successfully.")
			return
		}

		// If -f is not provided, check for existing config
		if err := ensureConfig(); err != nil {
			fmt.Println("Error:", err)
			if userWantsToCreateConfig() {
				if err := createConfigFile(); err != nil {
					fmt.Println("Error creating config:", err)
					os.Exit(1)
				}
			} else {
				fmt.Println("Initialization cancelled by user.")
				os.Exit(0)
			}
		} else {
			fmt.Println("Configuration verified successfully.")
		}
	},
}

func init() {
	initCmd.Flags().BoolVarP(&forceRecreate, "force", "f", false, "Force recreation of the configuration file")
	rootCmd.AddCommand(initCmd)
}
