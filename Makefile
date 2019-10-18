.RECIPEPREFIX += bifrost

BINARY_NAME := 
VERSION ?= latest
CURRENT_DIR := $(shell pwd)
PLATFORMS := linux windows macosx
os = $(word 1, $@)

all: build

build: linux windows macosx

clean:
  rm -rf build/*

run:
  go run $(BINARY_NAME).go

# Cross compilation
$(PLATFORMS):
  mkdir -p build
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o build/BINARY_NAME_$(VERSION)_$(os)_amd64 .