install:
	go mod tidy

lint:
	golangci-lint run