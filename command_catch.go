package main

import (
	"errors"
	"fmt"
	"math/rand"
	"pokedexcli/internal/pokeapi"
	"pokedexcli/internal/pokedex"
	"time"
)

func commandCatch(cfg *config, args ...string) error {
	if args == nil {
		return fmt.Errorf("Enter the name of a pokemon to try catch!")
	}

	pokemonName := args[0]
	pokemonInfoResp, err := cfg.pokeApiClient.GetPokemonInfo(pokemonName)
	if err != nil {
		var httpErr *pokeapi.ErrHTTP
		if errors.As(err, &httpErr) && httpErr.StatusCode == 404 {
			return fmt.Errorf(`Could not find Pokemon named "%v"`, pokemonInfoResp)
		}
		return err
	}

	newPokemon := pokedex.NewPokemonFromResp(pokemonInfoResp)

	fmt.Printf("Throwing a Pokeball at %v", newPokemon.Name)
	catchChance := rand.Intn(3 * newPokemon.BaseExp)
	catchRoll := rand.Intn(510)
	catchVal := catchChance - catchRoll
	var rollNumber int
	if catchVal < 0 {
		rollNumber = 4
	} else if catchVal < 20 {
		rollNumber = 3
	} else if catchVal < 40 {
		rollNumber = 2
	} else if catchVal < 70 {
		rollNumber = 1
	} else {
		rollNumber = 0
	}

	for i := 0; (i < rollNumber) && (i < 3); i++ {
		fmt.Print(".")
		time.Sleep(1500 * time.Millisecond)
	}

	switch rollNumber {
	case 4:
		fmt.Printf("!")
	case 3:
		fmt.Printf("\nShoot it was close! %v broke free.\n", newPokemon.Name)
	case 2:
		fmt.Printf("\nDang it! It looked like it was caught.\n")
	case 1:
		fmt.Printf("\n%v broke free of the PokeBall.\n", newPokemon.Name)
	default:
		fmt.Printf("\nThe PokeBall missed %v.\n", newPokemon.Name)
	}
	time.Sleep(500 * time.Millisecond)

	if rollNumber == 4 {
		fmt.Printf("\n%v was caught!\n", newPokemon.Name)
		cfg.pokedex.Add(newPokemon)
	}

	return nil

}
