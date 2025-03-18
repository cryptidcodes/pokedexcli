package main

import "fmt"

func commandExplore(cfg *config, locationName string) error {
	// takes a location name and fetches a list of pokemon located there
	fmt.Println("Exploring: ", locationName)

	// fetch the data
	locationAreaResp, err := cfg.pokeapiClient.ExploreLocation(locationName)
	if err != nil {
		fmt.Println("Error fetching data")
		return err
	}

	// print the pokemon names
	fmt.Println("Found Pokemon:")
	for _, pokemon := range locationAreaResp.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
