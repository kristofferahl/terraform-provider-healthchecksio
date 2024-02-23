package healthchecksio

import (
	"log"

	"github.com/kristofferahl/go-healthchecksio/v2"
)

const (
	// APIKeyEnvName contains the name of the environment variable to use when storing an api key
	APIKeyEnvName = "HEALTHCHECKSIO_API_KEY"
	APIURLEnvName = "HEALTHCHECKSIO_API_URL"
)

// Config contains healthchecksio provider settings
type Config struct {
	APIKey  string
	BaseURL string
}

// Client returns a configured healthchecksio client
func (c *Config) Client() (interface{}, error) {
	client := healthchecksio.NewClient(c.APIKey)
	if c.BaseURL != "" {
		client.BaseURL = c.BaseURL
	}

	log.Print("[INFO] healthchecks.io client configured")

	return client, nil
}
