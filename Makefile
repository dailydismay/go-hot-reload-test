SHELL := /bin/bash

tidy:
	go mod tidy
	go mod vendor

mocks:
	go generate ./..

test_handlers:
	go test gofirstapp/internal/handlers

dev:
	docker-compose up

dev_api:
	CompileDaemon --build="go build cmd/gotestapp/main.go" --command=./main
