package main

import (
	"fmt"
	"pokedexcli/internal/pokeapi"
)

func commandMap(c *config) error {
	resp, err := pokeapi.GetLocationData(c.Next)
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}

	c.Next = resp.Next
	if resp.Previous == nil {
		c.Previous = ""
	} else {
		c.Previous = *resp.Previous
	}

	if len(resp.Locations) == 0 {
		return fmt.Errorf("No locations returned")
	}

	for _, location := range resp.Locations {
		fmt.Printf("%v\n", location.Name)
	}

	return nil
}
