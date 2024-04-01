# System Pre-requisites

Go lang: 1.22.1
Docker
Docker-compose

## Contributing to the codebase

- Install required dependencies: `make install`
- Create an executable file: `make build`
- Run the executable file : `./airport-app-backend`
- Run the app locally without creating a build: `make run`
- Access the running app on [local](https://0.0.0.0:8080/)

## Run the test files
- Run the command `make test`

### Steps to setup postgres db locally using docker-compose

- Create a `.env` file in the project folder
- Add the following attributes in the `.env` file
     ```
        POSTGRES_USER=<postgres-username>
        POSTGRES_PASSWORD=<postgres-password>
        HOST=localhost
        DB_NAME=postgres
        PORT=5432
        SSL_MODE=disable
        
- Run `docker-compose -f docker-compose.yaml up`

## Steps to run Makefile

- run command `make all` in the terminal to clean, test and build in sequence.
- When editing the `Makefile` ensure to have appropriate indentation.
