# System Pre-requisites

Go lang: 1.22.1
Docker
Docker-compose

## Contributing to the codebase

- Install required dependencies: `go mod tidy`
- Create an executable file: `go build`
- Run the app locally: `go run main.go`
- Access the running app on [local](https://0.0.0.0:8080/)

### Steps to setup postgres db locally using docker-compose

- Create a `.env` file in the project folder
- Create `POSTGRES_USER=<postgres_username>` and `POSTGRES_PASSWORD=<postgres_password>` in that `.env` file
- Run `docker-compose -f docker-compose.yaml up`
