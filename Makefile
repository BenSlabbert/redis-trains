SHELL := /bin/bash

include .env
export

BIN="./bin"
SRC=$(shell find . -name "*.go")
GIT_COMMIT_ID = $(shell git log --format="%H" -n 1)
GIT_BRANCH_NAME = $(shell git symbolic-ref --short -q HEAD)

$(info BIN output dir: ${BIN})
$(info GIT_COMMIT_ID: ${GIT_COMMIT_ID})
$(info GIT_BRANCH_NAME: ${GIT_BRANCH_NAME})

.PHONY: all clean install_deps mod vet fmt proto-gen test build buildFast upx install_upx

default: all

all: clean mod vet fmt proto-gen test build upx

build:
	mkdir -p $(BIN)
	GOOS=linux ARCH=amd64 CGO_ENABLED=0 go build -tags netgo -a -v -ldflags "-s -w -X main.GitCommit=$(GIT_COMMIT_ID)" -o $(BIN)/cmd/archive cmd/archive/main.go
	GOOS=linux ARCH=amd64 CGO_ENABLED=0 go build -tags netgo -a -v -ldflags "-s -w -X main.GitCommit=$(GIT_COMMIT_ID)" -o $(BIN)/cmd/notification cmd/notification/main.go
	GOOS=linux ARCH=amd64 CGO_ENABLED=0 go build -tags netgo -a -v -ldflags "-s -w -X main.GitCommit=$(GIT_COMMIT_ID)" -o $(BIN)/cmd/sse cmd/sse/main.go
	GOOS=linux ARCH=amd64 CGO_ENABLED=0 go build -tags netgo -a -v -ldflags "-s -w -X main.GitCommit=$(GIT_COMMIT_ID)" -o $(BIN)/cmd/train cmd/train/main.go

buildFast:
	mkdir -p $(BIN)
	go build -ldflags "-s -w -X main.GitCommit=$(GIT_COMMIT_ID)" -o $(BIN)/cmd/archive.fast cmd/archive/main.go
	go build -ldflags "-s -w -X main.GitCommit=$(GIT_COMMIT_ID)" -o $(BIN)/cmd/notification.fast cmd/notification/main.go
	go build -ldflags "-s -w -X main.GitCommit=$(GIT_COMMIT_ID)" -o $(BIN)/cmd/sse.fast cmd/sse/main.go
	go build -ldflags "-s -w -X main.GitCommit=$(GIT_COMMIT_ID)" -o $(BIN)/cmd/train.fast cmd/train/main.go

# https://upx.github.io/
upx:
	upx -o $(BIN)/cmd/archive.upx $(BIN)/cmd/archive
	upx -o $(BIN)/cmd/notification.upx $(BIN)/cmd/notification
	upx -o $(BIN)/cmd/sse.upx $(BIN)/cmd/sse
	upx -o $(BIN)/cmd/train.upx $(BIN)/cmd/train

fmt:
	go fmt ./...
	@test -z $(shell gofmt -l $(SRC)) || (gofmt -d $(SRC); exit 1)

vet:
	go vet ./...
	staticcheck ./...

mod:
	go mod tidy
	go mod verify

test:
	go test ./...

proto-gen:
	protoc --go_out=gen/pb/train_stream_pb/ --go-grpc_out=gen/pb/train_stream_pb/ --go-grpc_opt=paths=source_relative pkg/stream/train_stream_v1.proto

clean:
	rm -rf $(BIN)

install_deps: install_upx
	go get -v ./...
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

install_upx:
	rm -rf /tmp/upx
	curl -L https://github.com/upx/upx/releases/download/v3.96/upx-3.96-amd64_linux.tar.xz -o /tmp/upx.tar.xz
	mkdir /tmp/upx
	tar xvf /tmp/upx.tar.xz -C /tmp/upx
	sudo mv /tmp/upx/upx-3.96-amd64_linux/upx /usr/local/bin/upx
	upx --version
