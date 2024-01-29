// Package apiserver runs the server
package apiserver

import (
	"log"
	"main/internal/app/esclient"
	"main/internal/app/server"
)

// Run creates a new Elasticsearch client and runs the server
func Run() {
	es, err := esclient.New()
	if err != nil {
		log.Fatal(err)
	}
	config := server.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	s := server.NewAPIServer(config, es)
	if err := s.Start(); err == nil {
		log.Fatal(err)
	}
}
