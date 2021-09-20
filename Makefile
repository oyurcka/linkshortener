.PHONY: build
build:
	go build -v ./cmd
	go build -v ./pkg

.DEFAULT_GOAL := build