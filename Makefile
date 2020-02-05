BINDIR = bin
CLIENT_BIN = $(BINDIR)/ct-cli
SERVER_BIN = $(BINDIR)/ct-server
VERSION := $(shell cat version.txt)


protos:
	protoc -I. -I$(GOPATH)/src  -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapi --go_out=plugins=grpc:. api/binance/*.proto

build-all: protos build-client build-server

test-build:
	go test -cover ./...

build-client:
	go build -ldflags="-X 'main.Version=$(VERSION)'" -o $(CLIENT_BIN) client/*.go

build-server:
	go build -ldflags="-X 'main.Version=$(VERSION)'"  -o $(SERVER_BIN) server/*.go
