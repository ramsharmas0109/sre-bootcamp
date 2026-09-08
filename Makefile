include .env
export

POSTGRESQL_URL = postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE)
TEST_POSTGRESQL_URL = postgres://postgres:dummy@localhost:5432/postgres?sslmode=disable

test_db_up:
	@echo "Starting test DB..."
	docker compose up -d test_db
	@echo "Waiting for test DB to be ready..."
	@for i in $$(seq 30); do docker compose exec -T test_db pg_isready -U postgres > /dev/null 2>&1 && exit 0; sleep 1; done; echo "Test DB did not become ready in time"; exit 1

test_db_down:
	@echo "Stopping test DB..."
	docker compose stop test_db

migrate_test_up:
	@echo "Running migrations on test DB..."
	migrate -database "$(TEST_POSTGRESQL_URL)" -path migrations up

migrate_test_down:
	@echo "Running migrations on test DB..."
	migrate -database "$(TEST_POSTGRESQL_URL)" -path migrations down

test: test_db_up migrate_test_up
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
