include .env

migrate-up:
	migrate -database ${DB_SOURCE} -path db/migrations up

migrate-down:
	migrate -database ${DB_SOURCE} -path db/migrations down

migrate-create:
	migrate create -dir db/migrations -ext sql -seq $(n)

down:
	docker-compose down && docker volume prune -f

up:
	docker-compose up -d

sqlc:
	rm -rf db/sqlc
	sqlc generate

start:
	go run main.go

.PHONY: migrate-up migrate-down migrate-create down up sqlc start