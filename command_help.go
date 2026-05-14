package main

import (
	"fmt"
	"strings"
)

func commandHelp(c *config, args ...string) error {
	var b strings.Builder
	fmt.Fprintf(&b, `
Welcome to the Pokedex!
Usage:

`)
	for _, cmd := range getCommands() {
		fmt.Fprintf(&b, "%v: %v\n", cmd.name, cmd.description)
	}
	fmt.Println(b.String())
	return nil
}
