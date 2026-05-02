package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationArea struct {
	// count    int    `json:"count"`
	Next      string  `json:"next"`
	Previous  *string `json:"previous"`
	Locations []struct {
		Name string `json:"name"`
		// url  string `json:"url"`
	} `json:"results"`
}

func GetLocationData(url string) (LocationArea, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationArea{}, fmt.Errorf("%w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, fmt.Errorf("%w", err)
	}

	var locationAreas LocationArea
	if err = json.Unmarshal(data, &locationAreas); err != nil {
		return LocationArea{}, fmt.Errorf("%w", err)
	}

	return locationAreas, nil
}
