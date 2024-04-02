package main

import (
	"cmd/app/app.go/internal/apiserver"
	"flag"
)

var tls bool

func init() {
	flag.BoolVar(&tls, "tls", false, "run tls connection")
	flag.Parse()
}

func main() {
	apiserver.Run(tls)
}
