#!/bin/sh

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o bookTrade .
docker build --rm -t book_trade .
