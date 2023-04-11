#!/usr/bin/env bash

export DB_HOST=localhost
export DB_PORT=5432
export DB_USERNAME=vinl
export DB_PASSWORD=test
export DB_NAME=vinl
export DB_SSLMODE=disable

systemctl start docker;
sleep 5;
docker stop postgres && docker rm postgres
docker run --name postgres -e POSTGRES_USERNAME=postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres

sleep 5;
psql -U postgres -h localhost -f ../../init.sql


go run .
