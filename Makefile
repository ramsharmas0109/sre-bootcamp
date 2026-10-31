# -include .env
# export

POSTGRESQL_URL = postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE)
TEST_POSTGRESQL_URL = postgres://$(TEST_DB_USER):$(TEST_DB_PASS)@$(TEST_DB_HOST):$(TEST_DB_PORT)/$(TEST_DB_NAME)?sslmode=$(SSL_MODE)

test_db_up:
	@echo "Starting test DB..."
	docker compose -f docker-compose.local.yaml up -d test_db

test_db_down:
	@echo "Stopping test DB..."
	docker compose stop test_db

migrate_test_up:
	@echo "Running migrations on test DB..."
	migrate -database "$(TEST_POSTGRESQL_URL)" -path migrations up

migrate_test_down:
	@echo "Running migrations on test DB..."
	migrate -database "$(TEST_POSTGRESQL_URL)" -path migrations down

configure_test_env: test_db_up migrate_test_up
	@echo "Configuring test environment..."

test:
	@echo "Running tests..."
	go test -v ./internal/handler

db_up:
	@echo "Starting dev DB..."
	docker compose up -d db

db_down:
	@echo "Stopping dev DB..."
	docker compose stop db

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
