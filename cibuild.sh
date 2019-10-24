#!/bin/bash
go get ./src/App
golangci-lint run ./src/App
go build ./src/App
