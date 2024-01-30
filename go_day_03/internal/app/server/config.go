package server

import (
	"main/internal/app/esclient"
)

const (
	minPage = 1
)

// Config is a struct used to store a port
type Config struct {
	Addr string
}

// NewConfig returns a new `Config` struct
func NewConfig() *Config {
	return &Config{
		Addr: ":8888",
	}
}

// JsonResponse is a struct representing a restaurant as JSON
type JsonResponse struct {
	Name     string           `json:"name"`
	Total    int              `json:"total,omitempty"`
	Places   []esclient.Place `json:"places"`
	PrevPage int              `json:"prev_page,omitempty"`
	NextPage int              `json:"next_page,omitempty"`
	LastPage int              `json:"last_page,omitempty"`
}

// NewJsonResponse returns a new `JsonResponse` struct
func NewJsonResponse(places []esclient.Place, counts int) *JsonResponse {
	return &JsonResponse{
		Name:   "Places",
		Total:  counts,
		Places: places,
	}
}

// NewRecommendationResponse returns a new `JsonResponse` struct for recommended restaurants
func NewRecommendationResponse(places []esclient.Place) *JsonResponse {
	return &JsonResponse{
		Name:   "Recommendation",
		Places: places,
	}
}

// Data is a struct for variables presented on a web page
type Data struct {
	Places   []esclient.Place
	Counts   int
	Page     int
	PrevPage int
	NextPage int
	LastPage int
}

// NewData returns a new `Data` struct
func NewData(places []esclient.Place, counts, page int) *Data {
	return &Data{
		Places:   places,
		Counts:   counts,
		Page:     page,
		PrevPage: page - 1,
		NextPage: page + 1,
		LastPage: counts/10 + 1,
	}
}
