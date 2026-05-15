package main

import (
	"fmt"
	"sort"
	"strings"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.pokedex.Caught) == 0 {
		return fmt.Errorf("Your Pokedex is empty.")
	}

	pokemonList := make([]string, 0, len(cfg.pokedex.Caught))
	for p, _ := range cfg.pokedex.Caught {
		pokemonList = append(pokemonList, p)
	}
	sort.Slice(pokemonList, func(i, j int) bool {
		return (cfg.pokedex.Caught[pokemonList[i]].Id < cfg.pokedex.Caught[pokemonList[j]].Id)
	})
	var b strings.Builder
	fmt.Fprintln(&b, "Your Pokedex:")
	for _, val := range pokemonList {
		fmt.Fprintf(&b, " - %v\n", val)
	}

	fmt.Print(&b)
	return nil

}
