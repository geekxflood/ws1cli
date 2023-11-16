// cmd/version.go

package cmd

import (
	"fmt"
	"geekxflood/ws1cli/internal/version"

	"github.com/spf13/cobra"
)

// versionCmd represents the 'version' command for printing the ws1cli version.
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of WS1CLI",
	Long:  `All software has versions. This is WS1CLI's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("WS1CLI version", version.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
