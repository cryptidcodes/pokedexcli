package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(locationName string) (RespLocationAreas, error) {
	url := baseURL + "/location-area"

	// append the location name to the url if it's not empty
	if locationName != "" {
		url += "/" + locationName
	}

	// check if the data is in the cache
	//fmt.Println("// checking cache for url:", url)
	cacheHit, ok := c.Cache.Get(url)

	// if the data is already cached, return it
	if ok {
		//fmt.Println("// cache hit, returning cache entry")
		locationAreaResp := RespLocationAreas{}
		//fmt.Println("// unmarshalling cache entry")
		err := json.Unmarshal(cacheHit, &locationAreaResp)
		if err != nil {
			return RespLocationAreas{}, err
		}
		return locationAreaResp, nil
	}

	// if not in the cache, fetch the data
	// fmt.Println("// cache miss, fetching data")
	// make the http request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationAreas{}, err
	}
	// send the request
	// fmt.Println("// sending request to: ", url)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationAreas{}, err
	}
	// close the response body when done
	defer resp.Body.Close()
	// read the response body to a byte slice (data)
	// fmt.Println("// reading response body")
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationAreas{}, err
	}
	// add the data to the cache
	// fmt.Println("// adding data to cache")
	c.Cache.Add(url, data)

	// create a new RespLocationAreas struct and unmarshal the data into it
	locationAreasResp := RespLocationAreas{}
	// fmt.Println("// unmarshalling data")
	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return RespLocationAreas{}, err
	}
	// return the struct and a nil error
	return locationAreasResp, nil
}
