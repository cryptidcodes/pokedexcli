package pokeapi

import (
	"net/http"
	"time"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

// Client is a struct that holds the http client
type Client struct {
	httpClient http.Client
	Cache      *pokeCache
}

// NewClient creates a new Client struct
func NewClient(timeout time.Duration) Client {
	pc := newPokeCache(30 * time.Second)
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		Cache: pc,
	}
}
