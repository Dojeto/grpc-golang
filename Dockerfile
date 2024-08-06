FROM golang:1.22 AS builder

WORKDIR /app

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

COPY . .

RUN go build -o server ./server/*.go

RUN go build -o client ./client/client.go

CMD ./server/handler & ./client/client