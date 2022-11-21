ifneq (,$(wildcard ./.env))
	include .env
	export
endif

postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER="${POSTGRES_USER}" -e POSTGRES_PASSWORD="${POSTGRES_PASSWORD}" -d postgres:15-alpine

createdb:
	docker exec -it postgres15 createdb --username="${POSTGRES_USER}" --owner="${POSTGRES_USER}" simple_bank

dropdb:
	docker exec -it postgres15 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://"${POSTGRES_USER}":"${POSTGRES_PASSWORD}"@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://"${POSTGRES_USER}":"${POSTGRES_PASSWORD}"@localhost:5432/simple_bank?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown