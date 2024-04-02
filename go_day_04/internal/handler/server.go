package handler

import (
	"cmd/app/app.go/internal/service"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Handler struct {
	Router *http.ServeMux
	Candy  *service.Service
}

func NewHandler() *Handler {
	h := &Handler{
		Router: http.NewServeMux(),
		Candy:  service.New(),
	}
	h.ConfigureRouter()
	return h
}

func (h *Handler) ConfigureRouter() {
	h.Router.HandleFunc("/buy_candy", h.BuyCandyHandler())
}

type Server struct {
	HTTPServer *http.Server
}

func NewHTTPServer(handler http.Handler, host, port string) *Server {
	return &Server{
		HTTPServer: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", host, port),
			Handler: handler,
		},
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Router.ServeHTTP(w, r)
}

func (s *Server) Start() {
	go func() {
		err := s.Run()
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()
	s.Stop(false)
}

func (s *Server) Run() error {
	log.Println("HTTP server is successfully started")
	return s.HTTPServer.ListenAndServe()
}

func (s *Server) Stop(f bool) {
	QuitChan := make(chan os.Signal)
	signal.Notify(QuitChan, syscall.SIGTERM, syscall.SIGINT)
	<-QuitChan
	if f {
		log.Println("HTTPS server is gracefully stopped")
	} else {
		log.Println("HTTP server is gracefully stopped")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := s.HTTPServer.Shutdown(ctx)
	if err != nil {
		log.Printf("Shutdown error: %v\n", err)
	}
}
