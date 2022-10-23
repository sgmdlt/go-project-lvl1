build:
	go build

install:
	go mod tidy

lint:
	golangci-lint run

run:
	./hexletcode