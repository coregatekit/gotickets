# App Configs
export APP_NAME=go-tickets
export APP_VERSION=1.0.0
export APP_PORT=8000
export APP_ENV=development

# Database Configs
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=antman
export DB_PASSWORD=3l}MSk1wg?7[
export DB_NAME=tickets

run:
	swag init -g main.go --parseDependency --output docs/
	go run main.go

build:
	rm -rf bin
	go build -o bin/main main.go

gen-openapi:
	swag init -g main.go --parseDependency --output docs/

test:
	go test -race -cover ./... -count=1 -failfast
	golangci-lint run

lint:
	golangci-lint run 