#!/usr/bin/env bash
# docker rm -f postgres
docker run -d --name postgres -e POSTGRES_PASSWORD=PASSWORD -e POSTGRES_DB=tg_bot -p 5432:5432 postgres:11
