SHELL := /bin/bash

.PHONY: setup fmt test project1-run kind-up kind-down

setup:
	@echo "Run setup guides in ./setup then execute: go mod tidy"

fmt:
	go fmt ./...

test:
	go test ./...

project1-run:
	go run ./projects/project-01-cli-yaml-auditor/cmd/auditor --input examples/manifests/insecure

kind-up:
	./kind/create-cluster.sh

kind-down:
	kind delete cluster --name go-sec-tooling
