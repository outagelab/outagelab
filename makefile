.PHONY: test docs

run:
	nv -- go run cmd/cli/main.go

dev-ui:
	cd ui && yarn dev

dev:
	nv -- npx concurrently --kill-others "make docs" "make run" "make dev-ui"

docs:
	cd docs && yarn && yarn dev

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
