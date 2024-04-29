# System Pre-requisites

- Go lang: 1.22.1
- Docker
- Docker-compose

## Steps to setup postgres db locally using docker-compose

- Create a `.env` file in the project folder and add the following entries

```bash
  export AIRPORT_POSTGRES_USER=<postgres-username>
  export AIRPORT_POSTGRES_PASSWORD=<postgres-password>
  export AIRPORT_HOST=localhost
  export AIRPORT_DB_NAME=<postgres-dbname>
  export AIRPORT_PORT=5432
  export AIRPORT_SSL_MODE=disable
```

## Contributing to the codebase

- To install required dependencies, run: `make install`
- To initiate the database, run: `make start-dependencies`
- To clean, test, and build app in sequence, run: `make all`
- To run all tests, run: `make test`
- To create an executable file, run: `make build`
- To run the packaged application, run: `./airport-app-backend`
- To run app without packaging, run: `make bootrun`
- The running app can be accessed [here](https://0.0.0.0:8080/)
- The Jaeger UI can be accessed [here](http://localhost:16686)
- To initiate swagger, run: `make swagger`
- The Swagger UI can be accessed [here](https://0.0.0.0:8080/swagger/index.html)

## Steps to setup SonarQube locally

- run `make start-dependencies`
- Hit [this link](http://localhost:9000/) [Credentials --> Username : Admin, Password : Admin]
- Create a local project named `airport` with project key as `airport` ![Image](Images/FirstStep.png)
- Change `main` to `master` ![Image](Images/LocalProject.png)
- Select an option of your choice for the second step ![Image](Images/SecondStep.png)
- Click `Create project`
- Select `locally` ![Image](Images/Locally.png)
- Give a token name of your choice and click `generate` (COPY THE TOKEN) ![Image](Images/Token.png)
- Select `other` for `Run analysis on your project` ![Image](Images/RunAnalysis.png)
- Install `sonar-scanner` using [homebrew](https://brew.sh/)
- Create a file named `sonar-project.properties` in the project root and paste the following lines (modify to your settings as appropriate)

```java
  sonar.projectKey=airport
  sonar.projectName=airport
  sonar.sources=.
  sonar.language=go
  sonar.sourceEncoding=UTF-8
  sonar.go.coverage.reportPaths=coverage.out
  sonar.coverage.exclusions=mocks/**,server/**,certs/**,config/**,database/**,**/*_test.go,**/main.go
  sonar.login=admin
  sonar.password=<your-password>
  sonar.token=<your-token>
```

- Run `make sonar-scan` to do analysis.
