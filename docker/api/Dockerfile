# TODO:change version
FROM golang:1.15-alpine3.13 AS build 
# ベースイメージの指定
ENV GO111MODULE=on
# 環境変数の設定

WORKDIR /go/src/github.com/ulwlu/go-rest
# 作業ディレクトリの指定。以降の,Dockerコンテナ内におけるカレントディレクトリはここで指定したものになる。

COPY ./src/go.mod ./src/go.sum ./
# modファイルをWORKDIRにCOPYする

RUN mkdir shell
COPY ./shell ./shell


RUN apk update \
    && apk add mariadb-client

RUN go mod download

COPY ./src ./ 

RUN pwd

EXPOSE 8081

CMD ["sh","./shell/start-server.sh"]
