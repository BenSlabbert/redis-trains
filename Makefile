.PHONY: build mod fmt vet proto-gen clean deps

build: clean mod fmt vet proto-gen
	go build -o bin/cmd/archive cmd/archive/main.go
	go build -o bin/cmd/notification cmd/notification/main.go
	go build -o bin/cmd/sse cmd/sse/main.go
	go build -o bin/cmd/train cmd/train/main.go

fmt:
	go fmt ./...

vet:
	go vet ./...

mod:
	go mod tidy
	go mod verify

proto-gen:
	protoc --go_out=gen/pb/train_stream_pb/ --go-grpc_out=gen/pb/train_stream_pb/ --go-grpc_opt=paths=source_relative pkg/stream/train_stream_v1.proto

clean:
	rm -rf bin/

deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
