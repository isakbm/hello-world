#!/bin/sh

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a \
    -tags netgo \
    -installsuffix netgo \
    -ldflags "-s -w" \
    -o ./server/server \
    ./server

docker build -t isakbm/hello-world-server:latest .

push() {
    docker push isakbm/hello-world-server:latest
}

push || sudo docker login && push

rm ./server/server