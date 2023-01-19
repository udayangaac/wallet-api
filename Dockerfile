FROM golang:1.19.2-alpine3.16

RUN mkdir /opt/wallet-api

ADD . /opt/wallet-api

WORKDIR /opt/wallet-api

RUN go build -o bin/wallet-api cmd/wallet-api/wallet_api.go 