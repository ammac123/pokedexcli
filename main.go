package main

import (
	"pokedexcli/internal/pokeapi"
	"pokedexcli/internal/pokedex"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	pokeDex := pokedex.NewPokedex()
	cfg := &config{
		pokeApiClient: pokeClient,
		pokedex:       pokeDex,
	}
	startRepl(cfg)
}
