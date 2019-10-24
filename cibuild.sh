#!/bin/bash
go get -d ./src/App
golangci-lint run ./src/App
go build ./src/App
