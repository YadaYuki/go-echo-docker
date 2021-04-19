# TODO:change version
FROM golang:1.15-alpine3.13 AS build

ENV GO111MODULE=on

WORKDIR /go/src/github.com/ulwlu/go-rest

COPY ./src/go.mod ./src/go.sum ./
RUN go mod download

COPY ./src ./ 
# ここで何やってんだ。。。？

EXPOSE 8081

CMD ["go","run","main.go"]
