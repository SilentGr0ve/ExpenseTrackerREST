include .env
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

migrate-