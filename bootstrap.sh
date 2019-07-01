#!/usr/bin/env bash
if [ -z "$1" ]; then
    echo "No storage path given!"
    exit 1
fi

mkdir -p $1

if [[ $? -ne 0 ]]; then
    echo "mkdir failed"
    exit 1
fi

docker stop docker-book-trade
docker rm -f docker-book-trade


docker run --name docker-book-trade -p 9000:8080 -d -v $1:/data/bookTrade book_trade
