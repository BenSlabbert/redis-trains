.PHONY: build fmt vet tidy

build: fmt vet
	go build -o bin/cmd/archive cmd/archive/main.go
	go build -o bin/cmd/notification cmd/notification/main.go
	go build -o bin/cmd/sse cmd/sse/main.go
	go build -o bin/cmd/train cmd/train/main.go

fmt:
	go fmt ./...

vet:
	go vet ./...

tidy:
	go mod tidy
