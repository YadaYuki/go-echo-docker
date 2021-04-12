# TODO:change version
FROM golang:1.15-alpine3.13 AS build

ENV GO111MODULE=on

WORKDIR /go/src/github.com/ulwlu/go-rest

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./ 
# ここで何やってんだ。。。？

CMD ["go","run","main.go"]
