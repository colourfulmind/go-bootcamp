.PHONY: build

build:
	go build -o tree cmd/tree/main.go

tests: build
	cd tests && go test -v

clean:
	rm -rf tree

.DEFAULT_GOAL := tests