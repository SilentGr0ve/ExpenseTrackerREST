-include .env
export

export PROJECT_ROOT=$(shell pwd)

env-up:
	@docker compose up -d tracker_postgres

env-down:
	@docker compose down tracker_postgres

env-port-forward:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder

migrate-up:
	@docker run --rm \
	--network host \
	-v $(PROJECT_ROOT)/migrations:/migrations \
	migrate/migrate:v4.19.1 \
	-path=/migrations \
	-database="postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(POSTGRES_DB)?sslmode=disable" \
	up

migrate-down:
	@docker run --rm \
	--netword host \
	-v $(POSTGRES_ROOT)/migrations:migrations \
	migrate/migrate:v4.19.1 \
	-path=/migrations \
	-database="postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(POSTGRES_DB)?sslmode=disable" \
	down 1

tracker-run:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	go mod tidy && \
	go run ${PROJECT_ROOT}/cmd/api/main.go
