run:
	go run main.go

build:
	rm -rf bin
	go build -o bin/main main.go

test:
	go test -race -cover ./... -count=1 -failfast
	golangci-lint run

lint:
	golangci-lint run 