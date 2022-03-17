SHELL := /bin/bash

generate-gql:
	go generate ./graph/

generate-wire:
	cd di && wire

generate: generate-gql generate-wire

run-local: generate
	go run cmd/todo/main.go server
