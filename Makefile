bootrun:
	make clean
	direnv allow
	make mock
	make test
	make test-report
	make docker
	make swagger
	make run
all:
	make clean
	make test
	make test-report
	make build
	make sonar
clean:
	go clean
	go mod tidy
run:
	go run main.go
test:
	go test ./...
test-report:
	mkdir -p build/reports/go-test-report && go test ./... -json | go-test-report -o build/reports/go-test-report/index.html
build:
	go build main.go
install:
	go get .
	go install github.com/vakenbolt/go-test-report@v0.9.3
	go install github.com/golang/mock/mockgen@v1.6.0
docker:
	docker-compose -f docker-compose.yaml up -d
mock:
	mockgen -destination=mocks/gate_service_mock.go -package=mocks airport-app-backend/services IGateRepository
	mockgen -destination=mocks/health_service_mock.go -package=mocks airport-app-backend/services IHealthRepository
	mockgen -destination=mocks/airline_service_mock.go -package=mocks airport-app-backend/services IAirlineRepository
sonar-scan:
	go test ./... -coverprofile=coverage.out
	sonar-scanner -X \
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
swagger:
	swag init
