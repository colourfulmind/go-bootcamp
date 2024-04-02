.PHONY: tests
tests:
	cd tests && go test -bench=. -cpuprofile=cpu.out -memprofile=mem.out
	cd tests && go tool pprof cpu.out > top10.txt

.PHONY: clean
clean:
	cd ./tests/ && rm -rf cpu.out mem.out tests.test