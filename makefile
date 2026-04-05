run:
    go run cmd/main.go

migrate-up:
    migrate -path migrations -database "$(DATABASE_URL)?sslmode=disable" up

migrate-down:
    migrate -path migrations -database "$(DATABASE_URL)?sslmode=disable" down

sqlc:
    sqlc generate

docker-up:
    docker compose up -d

docker-down:
    docker compose down