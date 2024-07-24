package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationArea() (LocationAreaResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
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

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	locAreaResp := LocationAreaResp{}
	err = json.Unmarshal(data, &locAreaResp)
	if err != nil {
		return LocationAreaResp{}, err
	}
	return locAreaResp, nil
}