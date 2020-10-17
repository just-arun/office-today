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
	./bin/office-today

connect:
	ssh root@128.199.30.69

deploy:
	make build
	scp $(PWD)/bin/office-today root@128.199.30.69:/root/server

push:
	scp $(PWD)/bin/office-today root@128.199.30.69:/root/server

clear:
	rm -rf ./main
	rm -rf ./bin/*
