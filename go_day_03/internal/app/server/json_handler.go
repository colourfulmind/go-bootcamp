package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"main/internal/app/db"
	"net/http"
)

// GetJsonHandler is a handler for `/api/places`
func (server *APIServer) GetJsonHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			places, counts, page := server.CreateDB(w, r)
			if page >= minPage && page <= db.GetCounts()/10+1 {
				data := NewJsonResponse(places, counts)
				data.FillPages(page)
				err := data.GetJson(w)
				if err != nil {
					ResponseError(w, err)
				}
			}
		} else {
			RequestError(w, errors.New(fmt.Sprintf("Method [%v] is not allowed", r.Method)))
		}
	}
}

// FillPages calculates the values for previous, next, and last pages
func (data *JsonResponse) FillPages(page int) {
	maxPage := db.GetCounts()/10 + 1
	if page > minPage {
		data.PrevPage = page - 1
	}
	if page < maxPage {
		data.NextPage = page + 1
	}
	if page != maxPage {
		data.LastPage = maxPage
	}
}

// GetJson outputs up to 10 restaurants depending on the page
func (data *JsonResponse) GetJson(w http.ResponseWriter) error {
	resp, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return errors.New("cannot marshall data")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(resp)
	return err
}
