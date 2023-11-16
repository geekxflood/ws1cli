// cmd/testFunctions.go

package cmd

import (
	"fmt"
	"io"
	"net/http"
)

// TestWS1 tests the configuration to Workspace ONE UEM.
func TestWS1() error {
	// Retrieve the configuration
	config, err := GetConfig()
	if err != nil {
		return fmt.Errorf("failed to get configuration: %w", err)
	}

	httpClient := CreateHTTPClient() // Use the global function to create an HTTP client

	// Construct the API endpoint URL from the configuration for system info
	apiEndpoint := fmt.Sprintf("%s%s/system/info", config.APIURL, config.APIPath)

	// Make a new HTTP request to the system info endpoint
	req, err := http.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return fmt.Errorf("creating the HTTP request failed: %w", err)
	}

	// Set the necessary headers for the request
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Basic "+config.DecryptedAPIAuth) // config.APIAuth is expected to be the base64 encoded 'username:password'

	if showDetails {
		fmt.Printf("URL: %s\n", apiEndpoint)
		fmt.Println("Headers:")
		for key, value := range req.Header {
			fmt.Printf("%v: %v\n", key, value)
		}
	}

	// Make the request using the created HTTP client
	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("making the HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	// Check the response status code to determine health
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body) // Read the body for additional error context
		return fmt.Errorf("API health check failed with status code: %d and body: %s", resp.StatusCode, string(body))
	}

	// Optionally, you could unmarshal and print the response body for user feedback
	fmt.Println("Workspace ONE UEM API is responsive.")
	return nil
}
