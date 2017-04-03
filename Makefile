# Default target: builds the project
build:
 go build -v -x ./cmd/temperature-server

# Cleans our project: deletes binaries
clean:
 go clean
