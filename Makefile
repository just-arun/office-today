setup:
	export GOBIN=$(pwd)/bin

dev:
	go run ./cmd/office-today/main.go

install:
	go install ./...

build:
	go build ./...

run:
	./main

clear:
	rm -rf ./main
	rm -rf ./bin/*