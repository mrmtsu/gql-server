SHELL := /bin/bash

install:
	go get github.com/99designs/gqlgen@v0.17.1

generate-gql:
	go generate ./graph/

generate-wire:
	cd di && wire

generate: install generate-gql generate-wire

run-local: generate
	go run cmd/todo/main.go server
