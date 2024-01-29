package server

import (
	"errors"
	"fmt"
	"html/template"
	"main/internal/app/esclient"
	"net/http"
	"strconv"
)

// GetPlacesHandler is a handler for the main page
func (server *APIServer) GetPlacesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			places, counts, page := server.CreateDB(w, r)
			if page != 0 {
				var res = NewData(places, counts, page)
				ts, err := template.ParseFiles("./web/template/index.html")
				if err != nil {
					ResponseError(w, err)
					return
				}
				err = ts.Execute(w, res)
				if err != nil {
					ResponseError(w, err)
				}
			}
		} else {
			RequestError(w, errors.New(fmt.Sprintf("Method [%v] is not allowed", r.Method)))
		}
	}
}

// GetPageNumber returns the current page number
func GetPageNumber(r *http.Request) (int, error) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if page == 0 {
		page = 1
		err = nil
	}
	return page, err
}

// CreateDB creates a database containing up to 10 restaurants depending on the current page
func (server *APIServer) CreateDB(w http.ResponseWriter, r *http.Request) ([]esclient.Place, int, int) {
	page, err := GetPageNumber(r)
	if err == nil && page >= minPage && page <= maxPage {
		places, counts, err := server.Store.Service.GetPlaces(page*10+1, page*10-9)
		if err != nil {
			ResponseError(w, err)
		}
		return places, counts, page
	} else {
		PageError(w, err, page)
	}
	return []esclient.Place{}, 0, 0
}
