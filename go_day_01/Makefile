EXEC = app

.PHONY: build
build:
	go build -o $(EXEC) ./cmd/app

.PHONY: run
run: build
	./$(EXEC) -f ./tests/database/original_database.xml
	./$(EXEC) -f ./tests/database/stolen_database.json
	./$(EXEC) --old ./tests/database/original_database.xml --new ./tests/database/stolen_database.json
	./$(EXEC) --old ./tests/database/snapshot1.txt --new ./tests/database/snapshot2.txt

.PHONY: test
test: build
	@cd tests && go test -v
	make clean

.PHONY: help
help: build
	./$(EXEC) -help

.PHONY: clean
clean:
	rm -rf $(EXEC)

.DEFAULT_GOAL := build