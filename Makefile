bootrun:
	make clean
	direnv allow
	make mock
	make test
	make docker
	make sonar
	make run
all:
	make clean
	make test
	make build
	make sonar
clean:
	go clean
	go mod tidy
run:
	go run main.go
test:
	mkdir -p build/reports/go-test-report && go test ./... -json | go-test-report -o build/reports/go-test-report/index.html
build:
	go build main.go
install:
	go get .
docker:
	docker-compose -f docker-compose.yaml up -d
mock:
	mockgen -destination=mocks/gate_service_mock.go -package=mocks airport-app-backend/services IGateRepository
	mockgen -destination=mocks/health_service_mock.go -package=mocks airport-app-backend/services IHealthRepository
	mockgen -destination=mocks/airline_service_mock.go -package=mocks airport-app-backend/services IAirlineRepository

sonar-scan:
	go test ./... -coverprofile=coverage.out
	sonar-scanner \
      -Dsonar.projectKey=Airport \
      -Dsonar.sources=. \
      -Dsonar.host.url=http://localhost:9000 \
      -Dsonar.token=$(SONAR_TOKEN)
sonar:
ifeq ($(CI),)
	make sonar-scan
else
	@echo "SonarQube scan skipped "
endif

