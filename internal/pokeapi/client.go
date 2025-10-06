package pokeapi

import (
	"net/http"
	"time"

	"github.com/b6sh/pokedexcli/internal/pokecache"
)

const baseUrl = "https://pokeapi.co/api/v2"

type Client struct {
	cache pokecache.Cache
	httpClient http.Client
}

func CreateClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client {
			Timeout: 10 * time.Second,
		},
	}
}