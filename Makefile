SHELL := /bin/bash

tidy:
	go mod tidy
	go mod vendor

run:
	go run ./cmd/gotestapp/main.go

mocks:
	go generate ./..

test_handlers:
	go test gofirstapp/internal/handlers
