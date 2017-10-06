#!/usr/bin/env bash
set -e
docker run -d -p 80:80 -v "$(pwd)"/:/data/:ro ezhai24/zipsite
go clean
