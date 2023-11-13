// cmd/initFunctions.go

package cmd

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
	"gopkg.in/yaml.v2"
)

func ensureConfig() error {
	if forceRecreate {
		return nil // Skip the check and proceed to recreate the config file
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user's home directory: %w", err)
	}
	configPath := homeDir + "/.ws1-cli.yaml"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return fmt.Errorf("configuration file does not exist at %s", configPath)
	}

	return nil
}

func userWantsToCreateConfig() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Would you like to create a new configuration file? (Y/n)")
	response, _ := reader.ReadString('\n')

	response = strings.TrimSpace(response)
	return response == "" || strings.ToLower(response) == "y"
}

func createConfigFile() error {
	var c Config
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter the WS1 URL (baseURL is enough):")
	c.APIURL, _ = reader.ReadString('\n')
	c.APIURL = strings.TrimSpace(c.APIURL)

	fmt.Println("Please enter the WS1 API Username:")
	ws1Username, _ := reader.ReadString('\n')
	ws1Username = strings.TrimSpace(ws1Username)

	fmt.Println("Enter WS1 API Password (input will be hidden):")
	ws1Password, err := term.ReadPassword(0)
	if err != nil {
		return fmt.Errorf("failed to read the password: %w", err)
	}

	// Combine the username and password encode the combined string in base64, then encrypt them
	combinedAuth := base64.StdEncoding.EncodeToString([]byte(ws1Username + ":" + string(ws1Password)))
	encryptedAuth, err := Encrypt([]byte(combinedAuth))
	if err != nil {
		return fmt.Errorf("failed to encrypt the auth credentials: %w", err)
	}
	c.APIAuth = encryptedAuth

	fmt.Println("\nEnter WS1 API Secret (input will be hidden):")
	ws1Secret, err := term.ReadPassword(0)
	if err != nil {
		return fmt.Errorf("failed to read the API secret: %w", err)
	}

	// Encrypt the API secret before storing it
	encryptedSecret, err := Encrypt(ws1Secret)
	if err != nil {
		return fmt.Errorf("failed to encrypt the API secret: %w", err)
	}
	c.APISecret = encryptedSecret

	fmt.Println("\nEnter WS1 API custom Path (leave blank for default '/API'):")
	c.APIPath, _ = reader.ReadString('\n')
	c.APIPath = strings.TrimSpace(c.APIPath)
	if c.APIPath == "" {
		c.APIPath = "/API"
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user's home directory: %w", err)
	}
	configPath := homeDir + "/.ws1-cli.yaml"

	// This will overwrite any existing file without needing to check for its existence
	yfile, err := yaml.Marshal(&c)
	if err != nil {
		return fmt.Errorf("failed to marshal the config: %w", err)
	}

	err = os.WriteFile(configPath, yfile, 0600) // Using os.WriteFile directly overwrites any existing file
	if err != nil {
		return fmt.Errorf("failed to write the config file at %s: %w", configPath, err)
	}

	fmt.Println("Configuration file created successfully at", configPath)
	return nil
}
