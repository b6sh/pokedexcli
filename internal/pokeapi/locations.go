package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)


type Locations struct {
	Count 	 int 	 `json:"count"`
	Next	 *string `json:"next"`
	Previous *string `json:"previous"`
	Results [] struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"results"`
}


func (c *Client) ListLocations (pageUrl *string) (Locations, error)  {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	if val, ok := c.cache.Get(url); ok {
		locations := Locations{}
		err := json.Unmarshal(val, &locations)
		if err != nil {
			return Locations{}, err
		}

		return locations, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Locations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Locations{}, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Locations{}, err
	}

	locations := Locations{}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return Locations{}, err
	}

	c.cache.Add(url, data)
	return locations, nil

}