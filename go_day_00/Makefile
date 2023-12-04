EXEC = app

.PHONY: build
build:
	go build -o $(EXEC) ./cmd/app

.PHONY: run
run: build
	./$(EXEC)

.PHONY: test
test:
	@cd tests && go test -v

.PHONY: test_files
test_files: build
	./$(EXEC) < ./tests/file1.txt
	./$(EXEC) < ./tests/file2.txt
	./$(EXEC) -mean -sd < ./tests/file1.txt
	./$(EXEC) -mode < ./tests/file1.txt
	./$(EXEC) -mode -median < ./tests/file1.txt
	./$(EXEC) -mode -median -mean < ./tests/file1.txt

.PHONY: clean
clean:
	rm -rf $(EXEC)

.PHONY: help
help: build
	./$(EXEC) -help

.DEFAULT_GOAL := build