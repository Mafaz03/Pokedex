package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationArea(pageURL *string) (LocationAreaResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}
	data, ok := c.cache.Get(fullURL)
	if ok {
		// cache hit
		fmt.Printf("Cache hit for %s\n",fullURL)
		locAreaResp := LocationAreaResp{}
		err := json.Unmarshal(data, &locAreaResp)
		if err != nil {
		return LocationAreaResp{}, err
	}
	return locAreaResp, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaResp{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResp{}, err
	}

	if resp.StatusCode > 399 {
		return LocationAreaResp{}, fmt.Errorf("error with resposnse status code: %v: %v", resp, resp.Status)
	}

	defer resp.Body.Close()

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	locAreaResp := LocationAreaResp{}
	err = json.Unmarshal(data, &locAreaResp)
	if err != nil {
		return LocationAreaResp{}, err
	}

	c.cache.Add(fullURL, data)

	return locAreaResp, nil
}



func (c *Client) ListPokeInLocation(locaiton string) (LocationExplore, error) {
	endpoint := "/location-area/" + locaiton
	fullURL := baseURL + endpoint

	data, ok := c.cache.Get(fullURL)
	if ok {
		// cache hit
		fmt.Printf("Cache hit for %s\n",fullURL)
		locationExplore := LocationExplore{}
		err := json.Unmarshal(data, &locationExplore)
		if err != nil {
		return LocationExplore{}, err
	}
	return locationExplore, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationExplore{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationExplore{}, err
	}

	if resp.StatusCode > 399 {
		return LocationExplore{}, fmt.Errorf("error with resposnse status code: %v: %v", resp, resp.Status)
	}

	defer resp.Body.Close()

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationExplore{}, err
	}

	locationExplore := LocationExplore{}
	err = json.Unmarshal(data, &locationExplore)
	if err != nil {
		return LocationExplore{}, err
	}

	c.cache.Add(fullURL, data)

	return locationExplore, nil
}