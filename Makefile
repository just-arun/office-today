setup:
	export GOBIN=$(PWD)/bin

dev:
	go run ./cmd/office-today/main.go

install:
	export GOBIN=$(PWD)/bin
	go install ./...

build:
	export GOBIN=$(PWD)/bin
	go build -o office-today ./cmd/office-today/main.go
	mv office-today ./bin

run:
	./main

clear:
	rm -rf ./main
	rm -rf ./bin/*