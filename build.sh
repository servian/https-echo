#!/bin/sh

docker build --pull -t vibrato/https-echo .
docker create --name https-echo vibrato/https-echo
docker cp https-echo:/https-echo .
docker rm https-echo
strip -u -x https-echo
