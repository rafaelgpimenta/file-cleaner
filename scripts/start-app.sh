#!/bin/bash

FOLDER="$( cd $(dirname "${BASH_SOURCE[0]}"); pwd )"

docker compose -f $FOLDER/../docker-compose/docker-compose.yml up -d

go run cmd/cleaner/main.go | jq .
