.PHONY: build
build:
	go build ./cmd/app/app.go

.PHONY: run
run: build
	./app

.PHONY: runTLS
runTLS: build
	./app -tls

.PHONY: run-client
run-client:
	go build ./cmd/client/candy-client.go
	./candy-client -k AA -c 2 -m 50

.PHONY: clean
clean:
	rm -rf app candy-client

.DEFAULT_GOAL := build
