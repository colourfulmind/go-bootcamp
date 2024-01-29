// Package store represents an interface to simply return the list of entries and enable pagination through them
package store

import (
	"github.com/elastic/go-elasticsearch/v8"
	"main/internal/app/db"
	"main/internal/app/esclient"
	//"net/http"
)

// Store is an interface designed to abstract the database
type Store interface {
	GetPlaces(limit int, offset int) ([]esclient.Place, int, error)
	GetClosestPlace(lat, lon string) ([]esclient.Place, error)
}

type ServeData struct {
	Service Store
}

// NewService returns a new ServeData struct
func NewService(es *elasticsearch.Client) *ServeData {
	return &ServeData{
		Service: db.Data{
			Es:     es,
			Places: []esclient.Place{},
			Counts: 0,
		},
	}
}
