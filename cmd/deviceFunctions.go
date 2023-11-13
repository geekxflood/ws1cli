// cmd/deviceFunctions.go

package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetDeviceInventory retrieves the inventory of devices based on the provided LGID.
func GetDeviceInventory(lgid int) error {
	config, err := GetConfig()
	if err != nil {
		return err
	}

	httpClient := CreateHTTPClient()

	const pageSize = 500
	var allDevices []DeviceDefinition
	var totalDevices int
	page := 0

	for {
		apiEndpoint := fmt.Sprintf("%s/API/mdm/devices/search?lgid=%d&page=%d", config.APIURL, lgid, page)

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

		var searchResult DeviceSearchResult
		bodyBytes, _ := io.ReadAll(resp.Body)

		if err := json.Unmarshal(bodyBytes, &searchResult); err != nil {
			return fmt.Errorf("error unmarshalling response: %v", err)
		}

		allDevices = append(allDevices, searchResult.Devices...)
		totalDevices = searchResult.Total

		// Break if the last page is reached
		if (page+1)*pageSize >= totalDevices {
			break
		}
		page++ // Go to the next page
	}

	// Convert all devices to JSON for a pretty print
	allDevicesJSON, err := json.MarshalIndent(allDevices, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling devices: %v", err)
	}
	fmt.Println(string(allDevicesJSON))

	return nil
}
