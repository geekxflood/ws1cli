// cmd/globalFunctions.go
package cmd

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"

	"gopkg.in/yaml.v2"
)

// CreateHTTPClient creates a new HTTP client, optionally skipping TLS verification.
func CreateHTTPClient() *http.Client {
	httpClient := &http.Client{}

	// Check if TLS verification should be skipped (insecure flag)
	if insecure {
		// Create a custom HTTP transport with TLSClientConfig
		httpClient.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	return httpClient
}

// HttpCaller makes an HTTP call to the specified API endpoint.
func HttpCaller(req *http.Request) (*http.Response, error) {
	httpClient := CreateHTTPClient()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API call failed with status code %d and body %s", resp.StatusCode, string(bodyBytes))
	}

	return resp, nil
}

// stringInSlice checks if a string is present in a slice.
func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

// GetConfig retrieves and decrypts the configuration.
func GetConfig() (*Config, error) {
	var c Config
	err := readConfig(&c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// readConfig reads the configuration file.
func readConfig(c *Config) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user's home directory: %w", err)
	}
	configPath := homeDir + "/.ws1-cli.yaml"
	return readYamlFile(c, configPath)
}

// readYamlFile reads and decodes the YAML configuration file.
func readYamlFile(c *Config, path string) error {
	yfile, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read configuration file: %w", err)
	}
	if err := yaml.Unmarshal(yfile, c); err != nil {
		return fmt.Errorf("failed to unmarshal configuration file: %w", err)
	}

	// Decrypt the APIAuth and APISecret after reading from the file
	decryptedAuth, err := DecryptValues(c.APIAuth)
	if err != nil {
		return fmt.Errorf("failed to decrypt APIAuth: %w", err)
	}
	decryptedSecret, err := DecryptValues(c.APISecret)
	if err != nil {
		return fmt.Errorf("failed to decrypt APISecret: %w", err)
	}

	// Store the decrypted values in the Config struct
	c.DecryptedAPIAuth = string(decryptedAuth)
	c.DecryptedAPISecret = string(decryptedSecret)

	return nil
}

// FilterDevices filters a JSON string of an array of DeviceDefinition.
func FilterDevices(devicesJSON string, valueFilter string) ([]string, error) {
	// Unmarshal the JSON string into a slice of DeviceDefinition
	var deviceDefinitions []DeviceDefinition
	err := json.Unmarshal([]byte(devicesJSON), &deviceDefinitions)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling devices: %v", err)
	}

	// Use reflection to loop over the fields of the struct and match the JSON tag.
	var filteredValues []string
	for _, device := range deviceDefinitions {
		val := reflect.ValueOf(device)
		typ := val.Type()

		for i := 0; i < val.NumField(); i++ {
			field := typ.Field(i)
			tag := field.Tag.Get("json")

			// If the tag matches the valueFilter, extract the value
			if tag == valueFilter {
				// Handle nested structs recursively if necessary.
				if field.Type.Kind() == reflect.Struct {
					nestedVal := val.Field(i)
					for j := 0; j < nestedVal.NumField(); j++ {
						nestedField := nestedVal.Type().Field(j)
						nestedTag := nestedField.Tag.Get("json")
						if nestedTag == valueFilter {
							filteredValues = append(filteredValues, fmt.Sprint(nestedVal.Field(j).Interface()))
						}
					}
				} else {
					filteredValues = append(filteredValues, fmt.Sprint(val.Field(i).Interface()))
				}
				break // Stop searching fields if we've found a match.
			}
		}
	}

	if len(filteredValues) == 0 {
		return nil, errors.New("no matching field found for the value filter provided")
	}

	return filteredValues, nil
}
