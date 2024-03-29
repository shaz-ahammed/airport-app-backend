# System Pre-requisites

Go lang: 1.22.1
Docker
Docker-compose

## Contributing to the codebase

- Install required dependencies: `go get .`
- Create an executable file: `go build`
- Run the executable file : `./airport-app-backend`
- Run the app locally without creating a build: `go run main.go`
- Access the running app on [local](https://0.0.0.0:8080/)

## Run the test files
- Go to folder that contains the test files. Eg: `cd middleware`
- Run the command `go test` to run all the test files present in that particular directory
- To run all the tests run, run the command `go test ./...`

### Steps to setup postgres db locally using docker-compose

- Create a `.env` file in the project folder
- Create `POSTGRES_USER=<postgres_username>` and `POSTGRES_PASSWORD=<postgres_password>` in that `.env` file
- Run `docker-compose -f docker-compose.yaml up`

## Steps to run Makefile

- run command `make` in the terminal to run all the commands in sequence listed in the `Makefile`
- When editing the `Makefile` ensure to have appropriate indentation.
