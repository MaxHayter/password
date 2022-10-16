#!/usr/bin/env bash
buf generate --template scripts/buf.gen.yaml --config scripts/buf.yaml --path api/proto/password.proto
