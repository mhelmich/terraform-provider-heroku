package heroku

import (
	"log"
	"net/http"
	"os"

	"github.com/heroku/heroku-go/v3"
)

type Config struct {
	Email   string
	APIKey  string
	Headers http.Header

	Api *heroku.Service
}

// Client returns a new Config for accessing Heroku.
func (c *Config) loadAndInitialize() error {
	var debugHTTP = false
	if os.Getenv("TF_LOG") == "TRACE" || os.Getenv("TF_LOG") == "DEBUG" {
		debugHTTP = true
	}
	c.Api = heroku.NewService(&http.Client{
		Transport: &heroku.Transport{
			Username:          c.Email,
			Password:          c.APIKey,
			UserAgent:         heroku.DefaultUserAgent,
			AdditionalHeaders: c.Headers,
			Debug:             debugHTTP,
		},
	})

	log.Printf("[INFO] Heroku Client configured for user: %s", c.Email)

	return nil
}
