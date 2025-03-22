package main

import (
	"fmt"
)

func commandInspect(cfg *config, pokemonName string) error {
	if pokemonName == "" {
		return fmt.Errorf("pokemon name required to use this function")
	}
	// retrieve the pokedex entry
	pokemon, ok := userPokedex.pokedexMap[pokemonName]
	if !ok {
		return fmt.Errorf("pokemon not caught yet")
	}
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)

	fmt.Println("Stats:")
	for i := 0; i < len(pokemon.Stats); i++ {
		fmt.Printf("-%v: %v\n", pokemon.Stats[i].Stat.Name, pokemon.Stats[i].BaseStat)
	}

	fmt.Println("Types:")
	for i := 0; i < len(pokemon.Types); i++ {

	}
	return nil
}
