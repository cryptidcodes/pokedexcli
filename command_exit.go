package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, arg string) error {
	// exits the program cleanly
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
