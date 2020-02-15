FROM golang:1.13
ENV PROTOC_ZIP=protoc-3.7.1-linux-x86_64.zip
ENV ENV_TAG dev
ENV GO_PROTO_TAG=v1.3.3
WORKDIR /app
ADD . /app
RUN apt update
RUN apt install unzip
RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/$PROTOC_ZIP
RUN unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
RUN unzip -o $PROTOC_ZIP -d /usr/local 'include/*'
ENV GO111MODULE=off
RUN  go get -d -u github.com/golang/protobuf/protoc-gen-go
RUN go env GOPATH
RUN ls /go/src/github.com
RUN git -C "$(go env GOPATH)"/src/github.com/golang/protobuf checkout $GO_PROTO_TAG
RUN go install github.com/golang/protobuf/protoc-gen-go
ENV GO111MODULE=on
EXPOSE 6000
RUN make ENV_TAG=${ENV_TAG} build-server-linux
ENTRYPOINT ["/app/bin/dev/ct-server"]