APP=anthro-env

.PHONY: build test fmt

build:
	go build -o bin/$(APP) ./cmd/$(APP)

test:
	go test ./...

fmt:
	gofmt -w ./cmd ./internal
