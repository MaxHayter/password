#!/usr/bin/env bash
# go get github.com/dizzyfool/genna
genna model-named -c "postgres://postgres:PASSWORD@localhost:5432/tg_bot?sslmode=disable" -o "internal/db/model/model.go" -t "public.*" -f -s deleted_at