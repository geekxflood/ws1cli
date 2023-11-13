// cmd/productFunctions.go

package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetProductInventory retrieves the inventory of products
func GetProductInventory() error {
	var jsonData []byte
	config, err := GetConfig()
	if err != nil {
		return err
	}
	httpClient := CreateHTTPClient()

	const pageSize = 500

	var allProducts []ProductDefinition
	var totalProducts int
	page := 0

	for {
		apiEndpoint := fmt.Sprintf("%s%s/mdm/products/search?lgid=%d", config.APIURL, config.APIPath, lgid)

		req, err := http.NewRequest("GET", apiEndpoint, nil)
		if err != nil {
			return err
		}

		req.Header.Add("Accept", "application/json")
		req.Header.Add("Authorization", "Basic "+config.DecryptedAPIAuth)
		req.Header.Add("aw-tenant-code", config.DecryptedAPISecret)

		resp, err := httpClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			bodyBytes, _ := io.ReadAll(resp.Body)
			return fmt.Errorf("API call failed with status code %d and body %s", resp.StatusCode, string(bodyBytes))
		}

		var searchResult ProductSearchResult
		bodyBytes, _ := io.ReadAll(resp.Body)

		if err := json.Unmarshal(bodyBytes, &searchResult); err != nil {
			return fmt.Errorf("error unmarshalling response: %v", err)
		}

		allProducts = append(allProducts, searchResult.Products...)
		totalProducts = searchResult.Total

		// Break if the last page is reached
		if (page+1)*pageSize >= totalProducts {
			break
		}
		page++ // Go to the next page
	}
	// Print all products
	if prettyPrint {
		jsonData, err = json.MarshalIndent(allProducts, "", "    ")
	} else {
		jsonData, err = json.Marshal(allProducts)
	}

	if err != nil {
		return fmt.Errorf("error marshalling devices: %v", err)
	}

	fmt.Println(string(jsonData))

	return nil
}
