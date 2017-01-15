#!/usr/bin/env bash

echo removing github.com/go-swagger/go-swagger from workspace
rm -rf $1/src/github.com/go-swagger/go-swagger
echo go getting github.com/go-swagger/go-swagger
go get -u github.com/go-swagger/go-swagger

echo checkout out the correct go-swagger revision: 75d539e0ea2e09636e4ba8dbb1afe62fe19e64d7
cd $1/src/github.com/go-swagger/go-swagger
git checkout 75d539e0ea2e09636e4ba8dbb1afe62fe19e64d7

go install github.com/go-swagger/go-swagger/cmd/swagger
