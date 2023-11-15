// cmd/deviceFunctions.go

package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetDeviceInventory retrieves the inventory of devices based on the provided LGID.
func GetDeviceInventory(lgid int) error {
	var jsonData []byte
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

		apiEndpoint := fmt.Sprintf("%s%s/mdm/devices/search?lgid=%d&page=%d", config.APIURL, config.APIPath, lgid, page)

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

	// Print all devices
	if prettyPrint {
		jsonData, err = json.MarshalIndent(allDevices, "", "    ")
	} else {
		jsonData, err = json.Marshal(allDevices)
	}

	if err != nil {
		return fmt.Errorf("error marshalling devices: %v", err)
	}

	fmt.Println(string(jsonData))

	return nil
}

// RunCommandOnDevices runs a command on a list of devices
func RunCommandOnDevices(command string, devicesFiltered []string, valueFilter string) error {
	config, err := GetConfig()
	if err != nil {
		return err
	}

	apiEndpoint := fmt.Sprintf("%s%s/mdm/devices/commands/bulk?command=%s&searchby=%s",
		config.APIURL, config.APIPath, command, valueFilter)

	payload := map[string]interface{}{
		"BulkValues": map[string][]string{
			"Value": devicesFiltered,
		},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshalling payload: %v", err)
	}

	req, err := http.NewRequest("POST", apiEndpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+config.DecryptedAPIAuth)
	req.Header.Add("aw-tenant-code", config.DecryptedAPISecret)

	resp, err := HttpCaller(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Construct the response JSON
	responseJSON := map[string]interface{}{
		"command":      command,
		"devicefilter": map[string][]string{valueFilter: devicesFiltered},
		"response":     resp.StatusCode,
	}

	// Marshal the response JSON
	responseData, err := json.Marshal(responseJSON)
	if err != nil {
		return fmt.Errorf("error marshalling response JSON: %v", err)
	}

	fmt.Println(string(responseData))
	return nil
}
