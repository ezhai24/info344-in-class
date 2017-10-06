#!/usr/bin/env bash
set -e
GOOS=linux go build
docker build -t ezhai24/zipsite .
docker push ezhai24/zipsite
go clean
