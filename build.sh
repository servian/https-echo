#!/bin/sh

docker build -t vibrato/https-echo .
docker create --name https-echo vibrato/https-echo
docker cp https-echo:/https-echo .
docker rm https-echo
