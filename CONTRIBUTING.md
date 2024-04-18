##Run Test
    -Navigate to middleware(cd middleware)
    -Use the following command to run the tests:
        go test

# System Pre-requisites

Go lang: 1.22.1
Docker
Docker-compose

### Steps to setup postgres db locally using docker-compose

- Create a `.env` file in the project folder
- Add the following attributes in the `.env` file
```
        export AIRPORT_POSTGRES_USER=<postgres-username>
        export AIRPORT_POSTGRES_PASSWORD=<postgres-password>
        export AIRPORT_HOST=localhost
        export AIRPORT_DB_NAME=<postgres-dbname>
        export AIRPORT_PORT=5432
        export AIRPORT_SSL_MODE=disable
```
## Contributing to the codebase

- To initiate the database: `make docker`
- To clean, test, initiate docker container and build app in sequence: `make all`
- Run the executable file : `./airport-app-backend`
- Access the running app on [local](https://0.0.0.0:8080/)

- Install required dependencies: `make install`
- Create an executable file: `make build`
- Run all tests: `make test`

- To run app without creating a build: `make bootrun`
- Access the running app on [local](https://0.0.0.0:8080/)

-Access the Jaeger UI in [local](http://localhost:16686)
