package server

import (
	"errors"
	"fmt"
	"main/internal/app/esclient"
	"net/http"
	"strings"
)

// GetClosestHandler is a handler for `/api/recommend`
func (server *APIServer) GetClosestHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorization := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authorization) != 2 {
			UnauthorizedError(w, errors.New("unauthorized token"))
		} else {
			token, err := AuthorizeToken(authorization[1])
			if err != nil {
				UnauthorizedError(w, errors.New("unauthorized token"))
			} else {
				if token.Valid {
					if r.Method == http.MethodGet {
						places, _, _ := server.CreateClosestDB(w, r)
						data := NewRecommendationResponse(places)
						err := data.GetJson(w)
						if err != nil {
							ResponseError(w, err)
						}
					} else {
						RequestError(w, errors.New(fmt.Sprintf("Method [%v] is not allowed", r.Method)))
					}
				} else {
					UnauthorizedError(w, errors.New("invalid token"))
				}
			}
		}
	}
}

// CreateClosestDB creates a database containing the 3 closest restaurants
func (server *APIServer) CreateClosestDB(w http.ResponseWriter, r *http.Request) ([]esclient.Place, int, int) {
	lon, lat := GetLonLat(r)
	places, err := server.Store.Service.GetClosestPlace(lon, lat)
	if err != nil {
		ResponseError(w, err)
	}
	return places, 0, 0
}

// GetLonLat extracts latitude and longitude
func GetLonLat(r *http.Request) (string, string) {
	lon := r.URL.Query().Get("lon")
	lat := r.URL.Query().Get("lat")
	return lat, lon
}
