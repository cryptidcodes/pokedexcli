package main

import (
	"time"

	"github.com/cryptidcodes/pokedexcli/internal/pokeapi"
)

var commands map[string]cliCommand

func main() {
	// initialize commands list and pokeapi client
	initCommands()
	pokeClient := pokeapi.NewClient(5 * time.Second)
	// create a pointer to a config struct
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	// start the REPL
	startRepl(cfg)
}

func initCommands() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays a the previous page of locations",
			callback:    commandMapb,
		},
	}
}
