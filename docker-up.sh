#!/usr/bin/env bash

#docker-compose -f docker-compose.yaml down
#docker-compose -f docker-compose-elk.yaml down
#docker-compose -f docker-compose-flink.yaml down

rm -f -r /Volumes/data/lark/*

docker-compose -f docker-compose.yaml up
docker-compose -f docker-compose-elk.yaml up
docker-compose -f docker-compose-flink.yaml up