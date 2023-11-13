// cmd/device.go

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var inventory bool
var lgid int

// deviceCmd represents the device command
var deviceCmd = &cobra.Command{
	Use:   "device",
	Short: "Interact with the devices managed by Workspace ONE UEM",
	Long: `Use the device command to perform actions related to devices within the Workspace ONE UEM platform.
For instance, you can retrieve an inventory of devices by specifying the --inventory flag along with a --lgid to denote the location group ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		if inventory {
			if lgid == 0 { // or any other check if LGID should not be zero
				fmt.Println("You must specify a location group ID with --lgid when using --inventory")
				return
			}
			err := GetDeviceInventory(lgid)
			if err != nil {
				fmt.Printf("Error retrieving device inventory: %v\n", err)
			}
		} else {
			fmt.Println("No action specified. Use the --inventory flag to get device inventory.")
		}
	},
}

func init() {
	deviceCmd.Flags().BoolVarP(&inventory, "inventory", "d", false, "Retrieve the list of devices inventory")
	deviceCmd.Flags().IntVarP(&lgid, "lgid", "l", 0, "Location Group ID is mandatory for retrieving inventory")
	rootCmd.AddCommand(deviceCmd)
}
