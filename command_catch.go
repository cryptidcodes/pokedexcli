package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, pokemonName string) error {
	if pokemonName == "" {
		return fmt.Errorf("pokemon name required to use this function")
	}

	// retrieve the pokemon data
	respPokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		fmt.Println("error fetching data")
		return err
	}

	// get the pokemon's base exp to determine catch rate
	baseExp := respPokemon.BaseExperience
	catchRate := (50.0 / float64(baseExp))
	attempt := rand.Float64()

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)
	if attempt < float64(catchRate) {
		fmt.Println("Caught", pokemonName, "!")
		// add the pokemon to the user's pokedex
		userPokedex.mu.Lock()
		defer userPokedex.mu.Unlock()
		userPokedex.pokedexMap[pokemonName] = respPokemon
		return nil
	} else {
		fmt.Println(pokemonName, "broke free!")
		return nil
	}
}
