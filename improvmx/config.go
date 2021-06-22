package improvmx

import (
	improvmx "github.com/issyl0/go-improvmx"
)

type Config struct {
	Token string
}

type Client struct {
	client *improvmx.Client
	config *Config
}

// Client returns a new client for accessing DNSimple.
func (c *Config) Client() (*Client, error) {
	client := improvmx.NewClient(c.Token)

	provider := &Client{
		client: client,
		config: c,
	}

	return provider, nil
}
