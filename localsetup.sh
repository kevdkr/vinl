#!/usr/bin/env bash

systemctl start docker;
sleep 5;
docker stop postgres && docker rm postgres
docker run --name postgres -e POSTGRES_USERNAME=postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres

sleep 5;
psql -U postgres -h localhost -f ./init.sql

go run .
