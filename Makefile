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
	go test ./...
build:
	go build main.go
install:
	go get .
