package apiserver

import (
	"cmd/app/app.go/internal/handler"
)

func Run(tls bool) {
	h := handler.NewHandler()
	if !tls {
		s := handler.NewHTTPServer(h, "localhost", "3333")
		s.Start()
	} else {
		s := handler.NewHTTPServer(h, "candy.tld", "3333")
		s.StartTLS()
	}
}
