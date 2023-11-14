// cmd/test.go

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test the configuration to your Workspace ONE UEM environment",
	Long:  `Test the configuration to your Workspace ONE UEM environment by making a call to the API.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test called")
		err := TestWS1()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	testCmd.Flags().BoolVarP(&showDetails, "show-details", "d", false, "Display the URL and headers for the test API call")
	rootCmd.AddCommand(testCmd)
}
