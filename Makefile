bootrun:
	direnv allow
	make clean
	make swagger
	make mock
	make install
	make start-dependencies
	make run
all:
	make clean
	make swagger
	make mock
	make install
	make test
	make build
clean:
	go clean
	go mod tidy
run:
	go run main.go
test:
	go test ./...
	mkdir -p build/reports/go-test-report && go test ./... -json | go-test-report -o build/reports/go-test-report/index.html
build:
	go build main.go
install:
	go get .
	go install github.com/vakenbolt/go-test-report@v0.9.3
start-dependencies:
	docker-compose -f docker-compose.yaml up -d
mock:
	go install github.com/golang/mock/mockgen@v1.6.0
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
	go install github.com/swaggo/swag/cmd/swag@v1.16.3
	swag init
