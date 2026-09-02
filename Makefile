POSTGRESQL_URL='postgres://postgres:ahl@123@localhost:5555/students?sslmode=disable'

migrate_up:
	@echo "Running migrations..."
	migrate -database ${POSTGRESQL_URL} -path migrations up

migrate_down:
	@echo "Running migrations..."
	migrate -database ${POSTGRESQL_URL} -path migrations down