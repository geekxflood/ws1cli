// cmd/remoteFunctions.go

package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetRemoteSession retrieves a remote session URL for a specified device and session type.
func GetRemoteSession(deviceUuid string, sessionType string) error {
	config, err := GetConfig()
	if err != nil {
		return err
	}

	// Construct the API endpoint URL.
	// Replace `apiPath` with the actual path where the remote session is created.
	apiEndpoint := fmt.Sprintf("%s%s/mdm/remote-management/devices/%s/session", config.APIURL, config.APIPath, deviceUuid)

	// Create the request body.
	requestBody, err := json.Marshal(map[string]string{
		"remote_management_tool_name": sessionType,
	})
	if err != nil {
		return err
	}

	// Create the request with the correct headers and body.
	req, err := http.NewRequest("POST", apiEndpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+config.DecryptedAPIAuth)
	req.Header.Add("aw-tenant-code", config.DecryptedAPISecret)

	// Make the request.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check the response.
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API call failed with status code %d and body %s", resp.StatusCode, string(body))
	}

	// Read and output the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("Remote session created successfully:")
	fmt.Println(string(body))
	return nil
}
