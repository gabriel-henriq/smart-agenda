include .env
.PHONY: migrate-up migrate-down migrate-create down up sqlc

migrate-up:
	migrate -database ${POSTGRESQL_URL} -path db/migrations up

migrate-down:
	migrate -database ${POSTGRESQL_URL} -path db/migrations down all

migrate-create:
	migrate create -dir db/migrations -ext sql -seq $(n)

down:
	docker-compose down && docker volume prune -f

up:
	docker-compose up -d

sqlc:
	rm -rf db/sqlc
	sqlc generate