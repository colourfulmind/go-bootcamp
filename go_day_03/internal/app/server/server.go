// Package server creates and starts the server
package server

import (
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"main/internal/app/db"
	"main/internal/app/store"
	"net/http"
)

// APIServer is a struct that contains data to start the API server correctly
type APIServer struct {
	Es     *elasticsearch.Client
	Router *http.ServeMux
	Store  *store.ServeData
	Config *Config
	Data   *db.Data
}

// NewAPIServer creates a new API server
func NewAPIServer(config *Config, es *elasticsearch.Client) *APIServer {
	s := &APIServer{
		Es:     es,
		Router: http.NewServeMux(),
		Store:  store.NewService(es),
		Config: config,
	}
	s.ConfigureRouter()
	return s
}

// ServeHTTP dispatches the request to the handler
func (server *APIServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	server.Router.ServeHTTP(w, r)
}

// ConfigureRouter configures handlers for the given pattern
func (server *APIServer) ConfigureRouter() {
	server.Router.HandleFunc("/", server.GetPlacesHandler())
	server.Router.HandleFunc("/api/places", server.GetJsonHandler())
	server.Router.HandleFunc("/api/recommend", server.GetClosestHandler())
	server.Router.HandleFunc("/api/get_token", server.GetTokenHandler())
}

// Start starts the API server
func (server *APIServer) Start() error {
	log.Println("Server is successfully started")
	return http.ListenAndServe(server.Config.Addr, server)
}
