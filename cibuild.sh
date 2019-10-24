#!/bin/bash
go get 
golangci-lint run .
go build ./cmd/Client
