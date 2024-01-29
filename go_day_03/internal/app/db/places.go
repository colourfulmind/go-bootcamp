package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"main/internal/app/esclient"
	"strings"
)

// Data contains all the needed data to make a request
type Data struct {
	Es     *elasticsearch.Client
	Places []esclient.Place
	Counts int
}

// ParseData parses data from a request and returns all restaurants
func ParseData(resp *esapi.Response) (esclient.Place, error) {
	var req Request
	err := json.NewDecoder(resp.Body).Decode(&req)
	if err != nil {
		return esclient.Place{}, err
	}
	var Place esclient.Place
	Place = *req.Hits.Hits[0].Source
	return Place, nil
}

// GetPlaces sends a request to get restaurants depending on the current page
func (d Data) GetPlaces(limit int, offset int) ([]esclient.Place, int, error) {
	for i := offset; i < limit; i++ {
		d.Counts = GetCounts()
		if limit > d.Counts {
			limit = d.Counts + 1
		}
		query := fmt.Sprintf(`{ "query": { "match": { "_id": "%d" } } }`, i)
		resp, err := d.Es.Search(
			d.Es.Search.WithIndex("places"),
			d.Es.Search.WithBody(strings.NewReader(query)),
			d.Es.Search.WithPretty(),
		)
		if err != nil {
			return nil, 0, errors.New(fmt.Sprintf("error : %s", err))
		}
		place, err := ParseData(resp)
		if err != nil {
			return nil, 0, err
		}
		if len(place.Name) != 0 {
			d.Places = append(d.Places, place)
		}
		err = resp.Body.Close()
		if err != nil {
			return nil, 0, err
		}
	}
	return d.Places, d.Counts, nil
}
