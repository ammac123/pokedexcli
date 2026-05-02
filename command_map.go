package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {
	locationsResp, err := cfg.pokeApiClient.ListLocations(cfg.nextLocationsUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationsUrl = locationsResp.Next
	cfg.previousLocationsUrl = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Printf("%v\n", location.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.previousLocationsUrl == nil {
		return errors.New("you're on the first page")
	}

	locationsResp, err := cfg.pokeApiClient.ListLocations(cfg.previousLocationsUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationsUrl = locationsResp.Next
	cfg.previousLocationsUrl = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Printf("%v\n", location.Name)
	}
	return nil
}
