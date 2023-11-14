package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)


// productCmd represents the product command
var productCmd = &cobra.Command{
	Use:   "product",
	Short: "Interact with products managed by Workspace ONE UEM",
	Long:  `Interact with products in Workspace ONE UEM, including starting or stopping products.`,
	Run: func(cmd *cobra.Command, args []string) {
		if lgid == 0 {
			fmt.Println("You must specify a location group ID with --lgid when using --inventory")
			return
		}

		if productID != 0 {
			// Logic to start or stop the product based on startStopProduct value
			err := StartStopProduct(productID, startStopProduct)
			if err != nil {
				fmt.Printf("Error in starting/stopping product: %v\n", err)
			}
			return
		}

		err := GetProductInventory()
		if err != nil {
			fmt.Printf("Error retrieving product list: %v\n", err)
		}
	},
}

func init() {
	productCmd.Flags().IntVarP(&lgid, "lgid", "l", 0, "Location Group ID is mandatory for retrieving inventory")
	productCmd.Flags().IntVarP(&productID, "product-id", "p", 0, "Product ID for starting or stopping a product")
	productCmd.Flags().BoolVarP(&startStopProduct, "start-stop", "s", false, "Flag to start (true/1) or stop (false/0) a product")
	rootCmd.AddCommand(productCmd)
}
