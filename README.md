# Marvel
Marvel experiment in Golang

## Prerequisites

Go must be installed

## Instructions

### Configuration
- Create a developer account at Marvel
- Create a folder **resources** under the root of the project
- Create the file **marvel-api.properties** under **resources**. That file should be in the following format and be filled with relevant data from Marvel Developer Account

```sh
# ****** DO NOT ADD THAT FILE IN VCS ********

# DO NOT CHANGE marvel.test. It is used for unit test reasons.
marvel.test=success
# API KEY (Marvel Developer Account)
marvel.apikey=abcdefghijklmn
# API HASH (Marvel Developer Account)
marvel.hash=abcdefghijklmn
```

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