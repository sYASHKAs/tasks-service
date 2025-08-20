DB_DSN := "postgres://postgres:postgres@localhost:5434/tasksdb?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

migrate:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

run:
	go run cmd/server/main.go

lint:
	golangci-lint run --color=always
