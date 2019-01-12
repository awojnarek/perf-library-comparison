GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=perftest

build:
	GOOS=linux  GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)_linux 
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)_macos 