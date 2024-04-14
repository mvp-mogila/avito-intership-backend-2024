POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
DB_HOST=127.0.0.1
DB_PORT=5432

DB_DSN="postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(DB_HOST):$(DB_PORT)/banners?sslmode=disable"

setup:
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.57.1

containers-up:
	docker compose up -d

containers-down:
	docker compose down

migrations-up:
	goose -dir db/migrations postgres $(DB_DSN) up

migrations-down:
	goose -dir db/migrations postgres $(DB_DSN) down

run: containers-up
	go run cmd/main/main.go



