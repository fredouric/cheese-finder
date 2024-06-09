package dataset

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/rs/zerolog/log"
)

const (
	base  = "https://public.opendatasoft.com/api/explore/v2.1/catalog/datasets/fromagescsv-fromagescsv/records?"
	limit = 10
)

func Fetch() ([]Cheese, error) {

	totalCount, err := getTotalCount()
	if err != nil {
		return nil, err
	}
	ch := make(chan []Cheese, (totalCount/limit)+1)
	var wg sync.WaitGroup

	for offset := 0; offset < totalCount; offset += limit {
		wg.Add(1)
		go func(offset int) {
			defer wg.Done()
			cheeses, err := getCheesesWithOffset(offset)
			if err != nil {
				log.Warn().Err(err).Msg("failed to get cheese")
				return
			}

			ch <- cheeses
		}(offset)
	}

	wg.Wait()
	close(ch)

	var allCheeses []Cheese
	for cheeses := range ch {
		allCheeses = append(allCheeses, cheeses...)
	}

	return allCheeses, nil
}

func getCheesesWithOffset(offset int) ([]Cheese, error) {
	url := fmt.Sprintf(base+"limit=%d&offset=%d", limit, offset)
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

	return response.Cheeses, nil
}

func getTotalCount() (int, error) {
	url := fmt.Sprintf(base+"limit=%d&offset=%d", 1, 0)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var response struct {
		TotalCount int `json:"total_count"`
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return 0, err
	}

	return response.TotalCount, nil
}
