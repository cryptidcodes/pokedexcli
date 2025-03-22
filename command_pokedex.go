package main

import (
	"fmt"
)

func commandPokedex(cfg *config, arg string) error {
	// lists all the names of all pokemone caught by the user
	if len(userPokedex.pokedexMap) == 0 {
		fmt.Println("You have caught 0 Pokemon. Go catch some!")
	}
	fmt.Printf("You have caught %v Pokemon:\n", len(userPokedex.pokedexMap))
	for _, pokemon := range userPokedex.pokedexMap {
		fmt.Println(pokemon.Name)
	}
	return nil
}
