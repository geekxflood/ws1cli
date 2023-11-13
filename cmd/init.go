/*
Copyright © 2023 Christophe Rime christopherime@me.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/term"
	"gopkg.in/yaml.v2"
)

type Config struct {
	APIURL    string `yaml:"apiurl"`
	APIAuth   string `yaml:"apiusername"`
	APISecret string `yaml:"apisecret"`
	APIPath   string `yaml:"apipath"`
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the CLI",
	Long: `This command sets up the CLI by guiding the user through the creation of a configuration file.
It will check for an existing config and prompt to create one if not found.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := ensureConfig(); err != nil {
			fmt.Println("Error:", err)
			if userWantsToCreateConfig() {
				if err := createConfigFile(); err != nil {
					fmt.Println("Error creating config:", err)
					os.Exit(1)
				}
			} else {
				fmt.Println("Initialization cancelled by user.")
				os.Exit(0)
			}
		} else {
			fmt.Println("Configuration verified successfully.")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func ensureConfig() error {
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

	// Combine the username and password, then encrypt them
	combinedAuth := ws1Username + ":" + string(ws1Password)
	encryptedAuth, err := Encrypt([]byte(combinedAuth), os.Getenv("WS1CLI_KEY"))
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
	encryptedSecret, err := Encrypt(ws1Secret, os.Getenv("WS1CLI_KEY"))
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

	yfile, err := yaml.Marshal(&c)
	if err != nil {
		return fmt.Errorf("failed to marshal the config: %w", err)
	}

	if err := os.WriteFile(configPath, yfile, 0600); err != nil {
		return fmt.Errorf("failed to write the config file at %s: %w", configPath, err)
	}

	fmt.Println("Configuration file created successfully at", configPath)
	return nil
}
