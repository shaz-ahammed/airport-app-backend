# System Pre-requisites

Go lang: 1.22.1
Docker
Docker-compose

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
``` 
- Create a `flyway.conf` file in `/flyway/conf`
- Add the following attributes in `flyway.conf`:
```
        flyway.url=jdbc:postgresql://127.0.0.1:5432/airport
        flyway.baselineOnMigrate=true
        flyway.user=<postgres-username>
        flyway.password=postgres-password>
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

