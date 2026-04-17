include .env
export

run:
	go run main.go
service-deploy:
	docker compose up -d application
service-undeploy:
	docker compose down application
migrate-up:
	migrate -path migrations -database $(PSQL) up
migrate-down:
	migrate -path migrations -database $(PSQL) down
migrate-clear:
	migrate -path migrations -database $(PSQL) force 000001
