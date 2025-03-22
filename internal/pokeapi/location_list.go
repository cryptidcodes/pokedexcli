package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// check if the data is in the cache
	fmt.Println("// checking cache for url:", url)
	cacheHit, ok := c.Cache.Get(url)
	if ok {
		fmt.Println("// cache hit, returning cache entry")
		locationsResp := RespLocations{}
		err := json.Unmarshal(cacheHit, &locationsResp)
		if err != nil {
			return RespLocations{}, err
		}
		return locationsResp, nil
	}

	// if not in the cache, fetch the data
	fmt.Println("// cache miss, fetching data")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocations{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocations{}, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocations{}, err
	}

	// add the data to the cache
	c.Cache.Add(url, data)

	locationsResp := RespLocations{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespLocations{}, err
	}

	// return the data struct
	return locationsResp, nil
}
