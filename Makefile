.PHONY: init run

init:
	docker-compose -f \scripts\docker-compose.yaml up -d

run:
	go run main.go

test:
	go test ./... -v

test-svc:
	go test ./service/. -v