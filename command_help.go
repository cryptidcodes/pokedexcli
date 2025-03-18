package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println("Displaying help message")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, cmd := range commands {
		fmt.Printf("%v: %v\n", cmd.name, cmd.description)
	}
	return nil
}
