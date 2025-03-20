# App Configs
export APP_NAME=go-tickets
export APP_VERSION=1.0.0
export APP_PORT=8000
export APP_ENV=development

# Database Configs
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=antman
export DB_PASSWORD=3lMSkQr1wg7
export DB_NAME=tickets

run:
	swag init -g cmd/api/main.go --parseDependency --output docs/
	go run cmd/api/main.go

build:
	rm -rf bin
	go build -o bin/main cmd/api/main.go

gen-openapi:
	swag init -g cmd/api/main.go --parseDependency --output docs/

test:
	go test -race -cover ./... -count=1 -failfast
	golangci-lint run

mock:
	mockery --dir=packages --all --recursive --output=tests/fakes  --outpkg=fakes
	mockery --dir=libs --all --recursive --output=tests/fakes  --outpkg=fakes
	mockery --dir=database --all --recursive --output=tests/fakes  --outpkg=fakes

lint:
	golangci-lint run 

# Migrations
dbConnectionString="postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"
# run with 'make migrate-create name=<migration_name>'
migrate-create:
	migrate create -ext sql -dir database/migrations -seq $(name)

migrate-up:
	migrate -database $(dbConnectionString) -path database/migrations up 1
migrate-down:
	migrate -database $(dbConnectionString) -path database/migrations down 1
# run with 'make migrate-up version=<number_of_migrations>'
migrate-up-to:
	migrate -database $(dbConnectionString) -path database/migrations up $(version)
# run with 'make migrate-down version=<number_of_migrations>'
migrate-down-to:
	migrate -database $(dbConnectionString) -path database/migrations down $(version)