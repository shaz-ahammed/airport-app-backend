bootrun:
	make clean
	direnv allow
	make mock
	make test
	make docker
	make run
all:
	make clean
	make test
	make build
clean:
	go clean
	go mod tidy
run:
	go run main.go
test:
	go test ./... -json | go-test-report 
build:
	go build main.go
install:
	go get .
docker:
	docker-compose -f docker-compose.yaml up -d
mock:
	mockgen -destination=mocks/gate_service_mock.go -package=mocks airport-app-backend/services IGateRepository
	mockgen -destination=mocks/health_service_mock.go -package=mocks airport-app-backend/services IHealthRepository
