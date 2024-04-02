all:
	make clean
	make test
	make run
clean:
	go clean
	go mod tidy
run:
	go run main.go
test:
	go test ./...
