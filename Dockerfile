FROM golang:1.21-alpine AS builder

RUN apk add --no-cache make protobuf protobuf-dev

WORKDIR /build

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

COPY go.mod go.sum ./

FROM alpine:latest

ENTRYPOINT ["./app"]