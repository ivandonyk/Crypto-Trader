ENV_TAG = local
BINDIR = bin/$(ENV_TAG)
CLIENT_BIN = $(BINDIR)/ct-cli
SERVER_BIN = $(BINDIR)/ct-server
VERSION := $(shell cat version.txt)
CT_SERVICE_NAME=ct-server
DOMAIN=gcr.io/bharrellcloud



protos:
	protoc -I. -I$(GOPATH)/src  -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapi --go_out=plugins=grpc:. api/binance/*.proto
	protoc -I. -I$(GOPATH)/src  -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapi --go_out=plugins=grpc:. api/coinbase/*.proto

build-all: protos build-client build-server-local

test-build:
	go test -cover ./...

build-client:
	go build -tags $(ENV_TAG) -ldflags="-X 'main.Version=$(VERSION)'" -o $(CLIENT_BIN) client/*.go

build-server-linux:
	GOOS=linux go build -tags $(ENV_TAG) -ldflags="-X 'main.Version=$(VERSION)'"  -o $(SERVER_BIN) server/*.go

build-server-local:
	GOOS=darwin go build -tags $(ENV_TAG) -ldflags="-X 'main.Version=$(VERSION)'"  -o $(SERVER_BIN) server/*.go

docker-build-server:
	docker build -t $(DOMAIN)/$(CT_SERVICE_NAME):$(VERSION) .

docker-publish-image:
	docker push $(DOMAIN)/$(CT_SERVICE_NAME):$(VERSION)

prep-manifest:
	mkdir -p tmp
	cp k8s/*.yml tmp/
	sed -i'.' s%CONTAINER%$(DOMAIN)/$(CT_SERVICE_NAME):$(VERSION)%g tmp/crypto_server.yml

apply-manifest:
	kubectl apply -f tmp/*.yml

cleanup:
	rm -rf tmp

.PHONY: docker-build-server docker-publish-image prep-manifest apply-manifest cleanup
deploy: docker-build-server docker-publish-image prep-manifest apply-manifest cleanup