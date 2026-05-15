package pokeapi

import (
	"encoding/json"
)

func (c *Client) ListLocations(pageURL *string) (RespLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	data, err := c.GetData(url)

	if err != nil {
		return RespLocations{}, nil
	}

	locationsResp := RespLocations{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespLocations{}, err
	}

	return locationsResp, nil
}

func (c *Client) GetLocationInfo(locationNameId string) (RespLocationInfo, error) {
	url := baseURL + "/location-area/" + locationNameId

	data, err := c.GetData(url)

	if err != nil {
		return RespLocationInfo{}, err
	}

	locationInfoResp := RespLocationInfo{}
	err = json.Unmarshal(data, &locationInfoResp)
	if err != nil {
		return RespLocationInfo{}, err
	}

	return locationInfoResp, nil

}
