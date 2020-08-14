package healthchecksio

import (
	"log"

	"github.com/icaho/go-healthchecksio"
)

const (
	// EnvironmentKey contains the name of the environment variable to use when storing an api key
	EnvironmentKey = "HEALTHCHECKSIO_API_KEY"
)

// Config contains healthchecksio provider settings
type Config struct {
	APIKey string
}

// Client returns a configured healthchecksio client
func (c *Config) Client() (interface{}, error) {
	client := healthchecksio.NewClient(c.APIKey)

	log.Print("[INFO] healthchecks.io client configured")

	return client, nil
}
