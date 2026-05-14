package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"pokedexcli/internal/pokecache"
	"time"
)

type ErrHTTP struct {
	StatusCode int
	Status     string
}

func (e *ErrHTTP) Error() string {
	return fmt.Sprintf("http error: %s", e.Status)
}

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient(timeout time.Duration, interval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(interval),
	}
}

func (c *Client) GetData(url string) ([]byte, error) {
	if cacheData, ok := c.cache.Get(url); ok {
		return cacheData, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		return nil, &ErrHTTP{StatusCode: res.StatusCode, Status: res.Status}
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	c.cache.Add(url, data)

	return data, nil
}
