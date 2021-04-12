FROM golang:1.16-alpine3.13 AS build

ENV GO111MODULE=on

WORKDIR /go/src/github.com/ulwlu/go-rest

COPY ./ ./


CMD ["go","run","main.go"]
