// Package db creates a database of restaurants
package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"main/internal/app/esclient"
	"strings"
)

// Request s a struct used to parse requests from the Elasticsearch client
type Request struct {
	Took    float64 `json:"took"`
	TimeOut bool    `json:"timed_out"`
	Shards  struct {
		Total      int64 `json:"total"`
		Successful int64 `json:"successful"`
		Skipped    int64 `json:"skipped"`
		Failed     int64 `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int64  `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string          `json:"_index"`
			Id     string          `json:"_id"`
			Score  float64         `json:"_score"`
			Source *esclient.Place `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

// ParseDataClosest parses data from a request and returns all restaurants
func ParseDataClosest(resp *esapi.Response) ([]esclient.Place, error) {
	var req Request
	err := json.NewDecoder(resp.Body).Decode(&req)
	if err != nil {
		return []esclient.Place{}, err
	}
	var Places []esclient.Place
	for _, hit := range req.Hits.Hits {
		Places = append(Places, *hit.Source)
	}
	return Places, nil
}

// GetClosestPlace sends a request to get the three closest restaurants
func (d Data) GetClosestPlace(lat, lon string) ([]esclient.Place, error) {
	query := `{
		"sort": [
			{
				"_geo_distance": {
				"location": {
				"lat": ` + lat + `,
				"lon": ` + lon + `
					},
				"order": "asc",
				"unit": "km",
				"mode": "min",
				"distance_type": "arc",
				"ignore_unmapped": true
				}
			}
		]
	}`
	resp, err := d.Es.Search(
		d.Es.Search.WithIndex("places"),
		d.Es.Search.WithBody(strings.NewReader(query)),
		d.Es.Search.WithPretty(),
		d.Es.Search.WithSize(3),
	)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error : %s", err))
	}
	d.Places, err = ParseDataClosest(resp)
	if err != nil {
		return nil, err
	}
	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return d.Places, nil
}
