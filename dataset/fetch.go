package dataset

import (
	"encoding/json"
	"net/http"
)

const url = "https://public.opendatasoft.com/api/explore/v2.1/catalog/datasets/fromagescsv-fromagescsv/records?limit=20"

func Fetch() ([]Cheese, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response struct {
		TotalCount uint16   `json:"total_count"`
		Cheeses    []Cheese `json:"results"`
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	var cheeses []Cheese
	cheeses = append(cheeses, response.Cheeses...)

	return cheeses, nil
}
