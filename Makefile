BINDIR = bin
MKBIN = mkdir -p bin
CLIENT_BIN = $(BINDIR)/cryto-trader
VERSION := $(shell cat version.txt)


build-all: build-client

make-bin:
	mkdir -p ./bin

build-client: make-bin
	go build -ldflags="-X 'main.Version=$(VERSION)'" -o $(CLIENT_BIN) client/*.go
