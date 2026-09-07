include .env
export

POSTGRESQL_URL = postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE)

migrate_up:
	@echo "Running migrations..."
	migrate -database "$(POSTGRESQL_URL)" -path migrations up

migrate_down:
	@echo "Running migrations..."
	migrate -database "$(POSTGRESQL_URL)" -path migrations down

build:
	@echo "Building the application..."
	go build -o bin/app main.go

run: build
	@echo "Running the application..."
	./bin/app

clean:
	@echo "Cleaning up..."
	rm -rf bin
