// cmd/types.go

package cmd

type Config struct {
	APIURL             string `yaml:"apiurl"`
	APIAuth            string `yaml:"apiusername"`
	APISecret          string `yaml:"apisecret"`
	APIPath            string `yaml:"apipath"`
	DecryptedAPIAuth   string `yaml:"-"`
	DecryptedAPISecret string `yaml:"-"`
}
