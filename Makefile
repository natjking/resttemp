all: build
   
get-deps:
	go get -t ./...

# Default target: builds the project
build:
	go build -v -x ./cmd/temperature-server

# Cleans our project: deletes binaries
clean:
	go clean
