#!/bin/sh
go test $1 -timeout=10000ms -coverprofile=$1/cover.out -short -v
go tool cover -func=$1/cover.out
rm -rf $1/cover.out
