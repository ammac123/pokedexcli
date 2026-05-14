package main

import (
	"errors"
	"fmt"
	"pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *config, args ...string) error {
	if args == nil {
		return fmt.Errorf("Please enter a location name.")
	}

	locationId := args[0]
	fmt.Printf("Exploring %v...\n", locationId)

	locationInfoResp, err := cfg.pokeApiClient.GetLocationInfo(locationId)
	if err != nil {
		var httpErr *pokeapi.ErrHTTP
		if errors.As(err, &httpErr) && httpErr.StatusCode == 404 {
			return fmt.Errorf(`Could not find location "%v"`, locationId)
		}
		return err
	}

	if len(locationInfoResp.PokemonEncounters) == 0 {
		fmt.Printf("No Pokemon found in %v\n", locationId)
		return nil
	}
	fmt.Printf("Found Pokemon:\n")
	for _, pokemon := range locationInfoResp.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
