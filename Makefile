GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=findarrlib
VERSION=$(shell cat ./VERSION)

all: build

.PHONY: build
build:
	rm -rf ./build/;
	mkdir ./build;
	$(MAKE) -s go-build

test:
	CGO_ENABLED=0 GOOS=linux $(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -rf ./build/

go-build:
	$(GOBUILD) -o ./build/$(BINARY_NAME)

vendor: 
	$(GOCMD) mod vendor
