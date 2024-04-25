.PHONY: build
build:
	go build cmd/blog/main.go
	./main --config=./config/config.yaml

.PHONY: build_client
build_client:
	go build -o client cmd/client/main.go
	./client --config=./config/config.yaml

.PHONY: clean
clean:
	rm -rf main client amazing_logo.png

generate_files:
	protoc -I protos/proto protos/proto/articles/articles.proto --go_out=./protos/gen/go --go_opt=paths=source_relative --go-grpc_out=./protos/gen/go --go-grpc_opt=paths=source_relative
