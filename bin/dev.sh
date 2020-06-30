#!/bin/sh
nodemon --exec go run ./src/main.go --signal SIGTERM
