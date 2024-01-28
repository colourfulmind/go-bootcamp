EXEC00 = myFind
EXEC01 = myWc
EXEC02 = myXargs
EXEC03 = myRotate

build:
	cd $(EXEC00)/cmd/app && go build -o ../../$(EXEC00)
	cd $(EXEC01)/cmd/app && go build -o ../../$(EXEC01)
	cd $(EXEC02)/cmd/app && go build -o ../../$(EXEC02)
	cd $(EXEC03)/cmd/app && go build -o ../../$(EXEC03)

tests: build
	cd tests && go test -v

clean:
	rm -rf $(EXEC00)/$(EXEC00)
	rm -rf $(EXEC01)/$(EXEC01)
	rm -rf $(EXEC02)/$(EXEC02)
	rm -rf $(EXEC03)/$(EXEC03)
	rm -rf ./tests/files/*tar.gz ./tests/*tar.gz

.DEFAULT_GOAL := build