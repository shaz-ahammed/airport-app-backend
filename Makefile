build:
	echo "executing build and test"
	go test ./...
	go run main.go
