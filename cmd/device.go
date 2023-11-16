// cmd/device.go

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deviceCmd represents the 'device' command.
var deviceCmd = &cobra.Command{
	Use:   "device",
	Short: "Interact with the devices managed by Workspace ONE UEM",
	Long: `Use the device command to perform actions related to devices within the Workspace ONE UEM platform.
For instance, you can retrieve an inventory of devices by specifying the --inventory flag along with a --lgid to denote the location group ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		if inventory {
			if lgid == 0 {
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
		if command != "" {
			// Check if the command is valid and part of the list of supported commands
			if !stringInSlice(command, commandTypes) {
				fmt.Printf("Invalid command specified. Supported commands are: %v\n", commandTypes)
				return
			}
			if inputJson == "" {
				fmt.Println("You must specify a json array of device IDs with --inputJson when using --command")
				return
			}
			if valueFilter != "" {
				fmt.Println("You cannot specify a value filter with --valueFilter when using --command")
				return
			}
			// Check if valueFilter is valid
			if !stringInSlice(valueFilter, valueFilterTypes) {
				fmt.Printf("Invalid value filter specified. Supported value filters are: %v\n", valueFilterTypes)
				return
			}
			filterDevices, err := FilterDevices(inputJson, valueFilter)
			if err != nil {
				fmt.Printf("Error filtering devices: %v\n", err)
				return
			}
			err = RunCommandOnDevices(command, filterDevices, valueFilter)
			if err != nil {
				fmt.Printf("Error running command on devices: %v\n", err)
				return
			}
		}
	},
}

func init() {
	deviceCmd.Flags().BoolVarP(&inventory, "inventory", "d", false, "Retrieve the list of devices inventory")
	deviceCmd.Flags().IntVarP(&lgid, "lgid", "l", 0, "Location Group ID is mandatory for retrieving inventory")
	deviceCmd.Flags().StringVarP(&command, "command", "c", "", "Command to run on the device")
	deviceCmd.Flags().StringVarP(&inputJson, "inputJson", "", "", "List of devices in an json array")
	deviceCmd.Flags().StringVarP(&valueFilter, "valueFilter", "f", "", "Filter the devices based on the value of a property")
	rootCmd.AddCommand(deviceCmd)
}
