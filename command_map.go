package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.next)
	if err != nil {
		fmt.Println("Error fetching data")
		return err
	}

	cfg.next = locationsResp.Next
	cfg.previous = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.previous == nil {
		return fmt.Errorf("you're on the first page already")
	}

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.previous)
	if err != nil {
		fmt.Println("Error fetching data")
		return err
	}

	cfg.next = locationsResp.Next
	cfg.previous = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
