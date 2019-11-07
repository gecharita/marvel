# Marvel
Marvel experiment in Golang

## Prerequisites

Go must be installed

## Instructions

### Configuration for LOCAL environment
- Create a developer account at Marvel
- Add the following env variables with your personal Marvel API credentials:
    1. **marvel_api_hash**=****
    2. **marvel_api_key**=****


### Run
```shell
go run main.go
```

## TEST

```sh
# For single package & verbose
go test . -v
# For all packages
go test ./...
```

## DOCKERIZE 

##### Initial setup
```sh
# Go modules for dependency management
go mod init marvel
```

##### Main script (single stage)
```sh
# Produces executable 'marvel'
go build
# Build the image
cd docker
docker build -t marvel .
# Run the images
docker run --name marvel-go-client -d -p 9090:9090 marvel
```

##### Main script (multiple stages)- produces smaller docker images

```sh
# Produces executable 'marvel'
go build
# Build the image
cd docker
docker build -t marvel-optimized -f multistage.Dockerfile .

# Run the images
docker run --name marvel-go-client-optimized -d -p 9090:9090 marvel-optimized
```

## CURL examples

```sh
curl localhost:9090/marvel/getall
```

## Wiki
```sh
# Prints all go envvariables
go env
```