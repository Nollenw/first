include .env
export

run:
	go run main.go
migrate-up:
	migrate -path migrations -database $(PSQL) up
migrate-down:
	migrate -path migrations -database $(PSQL) down
migrate-clear:
	migrate -path migrations -database $(PSQL) force 000001
