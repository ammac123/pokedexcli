package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi"
	"pokedexcli/internal/pokedex"
	"strings"
)

type config struct {
	pokeApiClient        pokeapi.Client
	nextLocationsUrl     *string
	previousLocationsUrl *string
	pokedex              pokedex.Pokedex
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		command, ok := getCommands()[commandName]

		var argsList []string
		for i := 1; i < len(words); i++ {
			argsList = append(argsList, words[i])
		}

		if ok {
			err := command.callback(cfg, argsList...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	textLower := strings.ToLower(text)
	output := strings.Fields(textLower)
	return output
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore [location]",
			description: "Get the pokemon found in the location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch [pokemon]",
			description: "Try to catch the pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect [pokemon]",
			description: "Inspect details of caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Show pokemon in your Pokedex",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
	}
}
