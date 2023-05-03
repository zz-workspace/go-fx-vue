#!/bin/sh
export $(cat .env) > /dev/null 2>&1;

docker network create -d overlay --attachable mynetwork

docker stack deploy -c docker-compose.yml traefik-proxy