dev:
	reflex -r '\.go$' -s -- sh -c "go run ./cmd/office-today/main.go"

install:
	go install ./cmd/office-today/main.go

build:
	go build ./cmd/office-today/main.go

run:
	./main

clear:
	rm -rf ./main