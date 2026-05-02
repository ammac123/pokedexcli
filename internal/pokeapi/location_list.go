package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespAreaLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespAreaLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespAreaLocations{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespAreaLocations{}, err
	}

	locationsResp := RespAreaLocations{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespAreaLocations{}, err
	}

	return locationsResp, nil
}
