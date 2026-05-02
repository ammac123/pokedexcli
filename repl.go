package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := 0
		fmt.Printf("Your command was: %s\n", commandName)
	}
}

func cleanInput(text string) []string {
	textLower := strings.ToLower(text)
	output := strings.Fields(textLower)
	return output
}
