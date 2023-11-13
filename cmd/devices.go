// cmd/devices.go

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getDevicesCmd represents the getDevices command
var getDevicesCmd = &cobra.Command{
	Use:   "getDevices",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getDevices called")
	},
}

func init() {
	rootCmd.AddCommand(getDevicesCmd)
}
