all: build build-image
   
get-deps:
	go get -t ./...

# Default target: builds the project
build:
	mkdir build
	CGO_ENABLED=O GOOS=linux go build -a -installsuffix cgo -o build/temperature-server -v -x ./cmd/temperature-server
build-image:
	cp DockerFile build/; cd build; docker build -t temperature-server -f DockerFile .

# Cleans our project: deletes binaries
clean:
	go clean
