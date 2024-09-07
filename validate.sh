#!/bin/zsh
   # Make sure we're in the project directory within our GOPATH
    go mod tidy
    go get 
    golangci-lint run
    go vet .
    go test .
