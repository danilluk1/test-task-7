services-up:
	docker compose -f docker-compose.dev.yml up

services-down:
	docker compose -f docker-compose.dev.yml down

migrateup:
	migrate -path internal/db/migration -database "postgresql://postgres:admin@localhost:5432/test_task_7?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/db/migration -database "postgresql://postgres:admin@localhost:5432/test_task_7?sslmode=disable" -verbose down
	
new_migration:
	migrate create -ext sql -dir internal/db/migration -seq $(name)

up:
	docker compose up

down:
	docker compose down -v

sqlc:
	sqlc generate --experimental

server:
	go run ./cmd/main.go

.PHONY: services-up services-down  migrate-create migrate-up sqlc up down server