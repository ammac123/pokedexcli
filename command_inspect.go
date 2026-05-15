package main

import (
	"fmt"
	"strings"
)

func commandInspect(cfg *config, args ...string) error {
	if args == nil {
		return fmt.Errorf("Enter the name of a caught pokemon to inspect")
	}

	pokemonName := args[0]
	pokemon, ok := cfg.pokedex.Caught[pokemonName]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon\n")
	}

	// String out
	var b strings.Builder
	fmt.Fprintf(&b, "Name: %v\n", pokemon.Name)
	fmt.Fprintf(&b, "Pokemon ID: %v\n", pokemon.Id)
	fmt.Fprintf(&b, "Height: %v\n", pokemon.Height)
	fmt.Fprintf(&b, "Weight: %v\n", pokemon.Weight)

	fmt.Fprintln(&b, "Stats:")
	fmt.Fprintf(&b, "  - HP: %v\n", pokemon.Stats.Hp)
	fmt.Fprintf(&b, "  - Attack: %v\n", pokemon.Stats.Attack)
	fmt.Fprintf(&b, "  - Defense: %v\n", pokemon.Stats.Defense)
	fmt.Fprintf(&b, "  - Special Attack: %v\n", pokemon.Stats.SpecialAttack)
	fmt.Fprintf(&b, "  - Special Defense: %v\n", pokemon.Stats.SpecialDefense)
	fmt.Fprintf(&b, "  - Speed: %v\n", pokemon.Stats.Speed)

	fmt.Fprintln(&b, "Types:")
	for _, t := range pokemon.Types {
		fmt.Fprintf(&b, " - %v\n", t)
	}

	fmt.Print(&b)

	return nil
}
