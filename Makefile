include .env
export

dev:
	go run main.go

goose_status:
	goose -dir schema postgres "host=$$DB_HOST user=$$DB_USER password=$$DB_PASSWORD dbname=$$DB_NAME port=$$DB_PORT sslmode=$$DB_SSL_MODE" status

goose_up:
	goose -dir schema postgres "host=$$DB_HOST user=$$DB_USER password=$$DB_PASSWORD dbname=$$DB_NAME port=$$DB_PORT sslmode=$$DB_SSL_MODE" up

sqlc_generate:
	sqlc generate
