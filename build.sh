#!/bin/sh

docker build -t vibrato/https-echo .
docker run --rm \
 -v ${PWD}:/go/src/github.com/vibrato/https-echo \
 -w /go/src/github.com/vibrato/https-echo \
 -e CGO_ENABLED=0 -e GOOS="linux" -e GOARCH="amd64" \
 golang:alpine go build -a -ldflags="-s -w" -o https-echo
