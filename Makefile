SHELL := /bin/bash

# The name of the executable (default is current directory name)
CURRENT_DIR := $(shell echo $${PWD\#\#*/})
TARGET := $(TARGET_PREFIX)$(CURRENT_DIR)$(TARGET_SUFFIX)
.DEFAULT_GOAL: $(TARGET)

# These will be provided to the target
VERSION := `git describe --tags --always --dirty=-dev`
BUILD := `git rev-parse HEAD`

# Use linker flags to provide version/build settings to the target
LDFLAGS=-ldflags "-X=gitlab.com/lorislab/mvnrc/cmd.version=$(VERSION) -X=gitlab.com/lorislab/mvnrc/cmd.build=$(BUILD)"

# go source files, ignore vendor directory
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: all build clean install uninstall fmt simplify check run

all: check install

$(TARGET): $(SRC)
	@go build $(LDFLAGS) -o $(TARGET)

build: $(TARGET)
	@true

clean:
	@rm -f $(TARGET)

install:
	@go install $(LDFLAGS)

uninstall: clean
	@rm -f $$(which ${TARGET})

fmt:
	@gofmt -l -w $(SRC)

simplify:
	@gofmt -s -l -w $(SRC)

check:
	@test -z $(shell gofmt -l main.go | tee /dev/stderr) || echo "[WARN] Fix formatting issues with 'make fmt'"
	@for d in $$(go list ./... | grep -v /vendor/); do golint $${d}; done
	@go tool vet ${SRC}

run: install
	@$(TARGET)