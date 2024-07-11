.PHONY: test

run:
	go run cmd/api/main.go

dev:
	cd ui && yarn dev

build:
	cd ui && yarn build-only

db:
	-psql -d postgres -c "CREATE DATABASE outagelab"
	go run cmd/db-migrate/main.go -dir ./db/migrations postgres "user=postgres dbname=outagelab sslmode=disable" up

db-down:
	go run cmd/db-migrate/main.go -dir ./db/migrations postgres "user=postgres dbname=outagelab sslmode=disable" reset

db-reset:
	make db-down
	make db

db-migration:
	cd db/migrations && goose create name sql
