DB_USER=root
DB_PASSWORD=root
DB_HOST=127.0.0.1
DB_PORT=5432
DB_NAME=payment
ENV=development
NAME=alter_table_transaction

migrate-new:
	sql-migrate new -env="${ENV}" -config="./config.yaml" $(NAME)

migrate-up:
	sql-migrate up -env="${ENV}" -config="./config.yaml" $(NAME)

run:
	go run cmd/main.go

.PHONY: migrate-new migration-up run