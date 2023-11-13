// cmd/product.go

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// productCmd represents the product command
var productCmd = &cobra.Command{
	Use:   "product",
	Short: "Interact with products managed by Workspace ONE UEM",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if lgid == 0 {
			fmt.Println("You must specify a location group ID with --lgid when using --inventory")
			return
		}
		err := GetProductInventory()
		if err != nil {
			fmt.Printf("Error retrieving product liost: %v\n", err)
		}
	},
}

func init() {
	productCmd.Flags().IntVarP(&lgid, "lgid", "l", 0, "Location Group ID is mandatory for retrieving inventory")
	rootCmd.AddCommand(productCmd)
}
