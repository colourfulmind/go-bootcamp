package handler

import (
	"crypto/tls"
	"errors"
	"log"
	"net/http"
)

const (
	CertFile = "cert/localhost/cert.pem"
	KeyFile  = "cert/localhost/key.pem"
)

func (s *Server) StartTLS() {
	go func() {
		err := s.RunTLS()
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTPS server ListenAndServe: %v", err)
		}
	}()
	s.Stop(true)
}

func (s *Server) RunTLS() error {
	log.Println("HTTPS server is successfully started")
	s.HTTPServer.TLSConfig, _ = GetTLSConfig()
	return s.HTTPServer.ListenAndServeTLS("", "")
}

func GetTLSConfig() (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(CertFile, KeyFile)
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	return config, nil
}
