#!/usr/bin/env bash
set -e
GOOS=linux go build
docker build -t ezhai24/testsite .
docker push ezhai24/testsite
go clean
