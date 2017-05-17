#!/bin/sh
go test ./... -timeout=10000ms -coverprofile=cover.out -short -v
go tool cover -func=cover.out
rm -rf cover.out
