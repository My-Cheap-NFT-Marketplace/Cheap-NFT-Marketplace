include .env

migrateup:
	migrate -path ./infrastructure/database -database "postgres://$(MARKETPLACE_USER):$(MARKETPLACE_PASSWORD)@$(MARKETPLACE_HOST):5432/marketplace?sslmode=disable" -verbose up

migratedown:
	migrate -path ./infrastructure/database -database "postgres://$(MARKETPLACE_USER):$(MARKETPLACE_PASSWORD)@$(MARKETPLACE_HOST):5432/marketplace?sslmode=disable" -verbose down

compose:
	docker-compose up -d