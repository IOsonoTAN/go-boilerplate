#!/bin/sh
go mod vendor
go build -o ./cmd/room-reservation ./src/main.go
./cmd/room-reservation
