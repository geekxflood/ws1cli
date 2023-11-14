// cmd/globalFunctions.go
package cmd

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

// CreateHTTPClient creates a new HTTP client based on the insecure flag.
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

// HttpCaller makes an HTTP call to the specified API endpoint
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

// GetConfig retrieves and decrypts the configuration from the config file
func GetConfig() (*Config, error) {
	var c Config
	err := readConfig(&c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// readConfig reads the configuration file and fills the Config struct
func readConfig(c *Config) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user's home directory: %w", err)
	}
	configPath := homeDir + "/.ws1-cli.yaml"
	return readYamlFile(c, configPath)
}

// readYamlFile reads and decodes the YAML configuration file
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
