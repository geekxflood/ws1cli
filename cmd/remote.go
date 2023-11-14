// cmd/remote.go

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// remoteCmd represents the remote command
var remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "Interact with remote mangement by Workspace ONE UEM",
	Long: ` Interact with remote mangement in Workspace ONE UEM,
	define the type of session to be created, and the device to be targeted.
	The command return a session URL that can be used to connect to the device.`,
	Run: func(cmd *cobra.Command, args []string) {
		if deviceUuid == "" {
			fmt.Println("You must specify a device UUID with --deviceUuid")
			return
		}
		if sessionType == "" {
			fmt.Println("You must specify a session type with --sessionType")
			return
		}
		// check if the sessionType is a value in sessionTypes
		found := false
		for _, t := range sessionTypes {
			if t == sessionType {
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("Invalid session type %s\n", sessionType)
			return
		}

		err := GetRemoteSession(deviceUuid, sessionType)
		if err != nil {
			fmt.Printf("Error retrieving remote session: %v\n", err)
		}
	},
}

func init() {
	remoteCmd.Flags().StringVarP(&deviceUuid, "deviceUuid", "u", "", "Device UUID")
	remoteCmd.Flags().StringVarP(&sessionType, "sessionType", "s", "", "type of session to be created 'ScreenShare', 'FileManager', 'RemoteShell', 'RegistryEditor'")
	rootCmd.AddCommand(remoteCmd)
}
