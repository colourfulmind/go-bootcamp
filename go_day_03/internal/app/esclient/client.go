// Package esclient creates a new Elasticsearch client
package esclient

import (
	"github.com/elastic/go-elasticsearch/v8"
)

const (
	address      = "http://localhost:9200"
	databasePath = "./assets/data.csv"
	del          = '\t'
	index        = "places"
	jsonPath     = "./api/schema.json"
)

// New creates and returns a new Elasticsearch client
func New() (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			address,
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	CheckIndex(es)
	return es, nil
}
