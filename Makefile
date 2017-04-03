all: build build-image
   
get-deps:
	go get -t ./...

# Default target: builds the project
build:
	CGO_ENABLED=O GOOS=linux go build -a -installsuffix cgo -v -x ./cmd/temperature-server
build-image:
	docker build -t temperature-server -f Dockerfile .

# Cleans our project: deletes binaries
clean:
	go clean
