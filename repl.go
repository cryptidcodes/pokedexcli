package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/cryptidcodes/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	next          *string
	previous      *string
}

func startRepl(cfg *config) {
	// creatte a scanner to read input from the user
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		// read a line from the input
		scanned := scanner.Scan()
		if !scanned {
			break
		}

		// get the input text
		var input []string = cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		// if there is no argument, append an empty string
		if len(input) == 1 {
			input = append(input, "")
		}

		// check if the command exists in the commands map
		command, exists := commands[input[0]]
		if exists {
			err := command.callback(cfg, input[1])
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
	// split input into words based on whitespace
	// lowercase and trim whitespace from prefix/suffix

	newSlice := []string{}
	tempString := ""

	for _, c := range text {
		if unicode.IsSpace(c) {
			if len(tempString) > 0 {
				newSlice = append(newSlice, strings.ToLower(tempString))
			}
			tempString = ""
		} else {
			tempString += string(c)
		}
	}
	if tempString != "" {
		newSlice = append(newSlice, strings.ToLower((tempString)))
	}
	return newSlice
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}
