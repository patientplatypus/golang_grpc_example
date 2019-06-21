#!/bin/bash

./kill.sh
# ./dockerBuildPush.sh

# rm -rf ./main
# go build main.go

docker-compose build
docker-compose up