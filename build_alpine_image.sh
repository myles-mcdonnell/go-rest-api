#!/usr/bin/env bash

docker run -t -i --name build_temp -v `pwd`:/go/src/github.com/myles-mcdonnell/go-rest-api golang:1.7.1-alpine go install github.com/myles-mcdonnell/go-rest-api/cmd/go-rest-server
docker cp build_temp:/go/bin/go-rest-server ./go-rest-server
docker rm build_temp
docker build -t go-rest-api .