FROM golang:1.22

WORKDIR /app

COPY proto proto

RUN apt-get update \
 && DEBIAN_FRONTEND=noninteractive \
    apt-get install --no-install-recommends --assume-yes \
      protobuf-compiler

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


RUN protoc --go_out=. --go-grpc_out=. proto/*.proto