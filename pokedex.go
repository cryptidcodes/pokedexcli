package main

import (
	"sync"

	"github.com/cryptidcodes/pokedexcli/internal/pokeapi"
)

type Pokedex struct {
	pokedexMap map[string]pokeapi.RespPokemon
	mu         sync.Mutex
}
