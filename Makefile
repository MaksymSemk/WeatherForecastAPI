include .env

MIGRATE_CONTAINER=migrate/migrate
MIGRATIONS_DIR=$(CURDIR)/cmd/migrate/migrations
DB_URL=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@host.docker.internal:5432/$(POSTGRES_DB)?sslmode=disable

.PHONY: up migrate create

run:
	docker compose up --build backend

up:
	docker-compose up --build
 
migrate-up:
	docker compose up migrate-up
migrate-down:
	docker compose up migrate-down

create:
	docker run --rm -v "$(MIGRATIONS_DIR):/migrations" $(MIGRATE_CONTAINER) \
		create -ext sql -dir /migrations -seq $(name)

clean-containers:
	docker-compose down -v