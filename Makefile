.PHONY: run test

run:
	go run main.go

test:
	go clean -testcache
	go test ./service/. -v