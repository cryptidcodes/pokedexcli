package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (RespPokemon, error) {
	if pokemonName == "" {
		return RespPokemon{}, fmt.Errorf("pokemon name required to use this function")
	}
	url := baseURL + "/pokemon/" + pokemonName

	// check if the data is in the cache
	cacheHit, ok := c.Cache.Get(url)
	if ok {
		pokemonResp := RespPokemon{}
		err := json.Unmarshal(cacheHit, &pokemonResp)
		if err != nil {
			return RespPokemon{}, err
		}
		return pokemonResp, nil
	}

	// if not in the cache, fetch the data
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemon{}, err
	}

	// add the data to the cache
	c.Cache.Add(url, data)

	pokemonResp := RespPokemon{}
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return RespPokemon{}, err
	}

	// return the data struct
	return pokemonResp, nil
}
