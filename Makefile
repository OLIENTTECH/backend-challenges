DB_HOST=localhost
DB_PORT=5432
DB_NAME=postgres
DB_USERNAME=postgres
DB_PASSWORD=P@ssw0rd
SSL_MODE=disable
DB_STRING := postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE) ## DBへの接続先

.PHONY: build
build:
	docker-compose -f docker-compose.dev.yml build --no-cache
	docker-compose -f docker-compose.dev.yml up -d

.PHONY: up
up:
	docker-compose -f docker-compose.dev.yml up -d

.PHONY: down
down:
	docker-compose -f docker-compose.dev.yml down


# ==========================
# Migration
# ==========================

.PHONY: migrate-up
migrate-up: ## Create tables to the latest version
	docker compose -f docker-compose.dev.yml exec -T api sh -c "goose -dir=schema/ddl -allow-missing up"

.PHONY: migrate-seed
migrate-seed: ## Seed mock data
	docker compose -f docker-compose.dev.yml exec -T api sh -c "goose -dir=schema/develop/dml -allow-missing up"

.PHONY: migrate-reset
migrate-reset: ## Reset tables and seed
	make migrate-drop
	make migrate-up

.PHONY: migrate-drop
migrate-drop: ## Rollback all migrations
	docker compose -f docker-compose.dev.yml exec -T api sh -c "goose -dir=schema/develop/dml -allow-missing reset && goose -dir=schema/ddl reset"

.PHONY: create-ddl
create-ddl: ## Create migration files
	docker-compose -f docker-compose.dev.yml exec api goose -dir=schema/ddl create create_${FILE} sql

.PHONY: create-dml
create-dml: ## Create migration files
	docker-compose -f docker-compose.dev.yml exec api goose -dir=schema/develop/dml create insert_${FILE} sql

# ==========================
# Test
# ==========================

.PHONY: lint
lint: ## Run lint
	docker compose -f docker-compose.dev.yml exec -T api sh -c "golangci-lint run"

.PHONY: test
test: ## Run test
	docker compose -f docker-compose.dev.yml exec -T api go test ./backend/...

# ==========================
# Build
# ==========================

.PHONY: build-hotreload
build-hotreload: ## Build hotreload
	GOOS=linux GOARCH=amd64 go build -tags timetzdata -o ./tmp ./cmd/*

.PHONY: build-api
build-api: ## Build api
	GOOS=linux GOARCH=amd64 go build -tags timetzdata -o ./bin ./cmd/*

# ==========================
# mock
# ==========================

.PHONY: mockgen
mockgen:
	docker-compose -f docker-compose.dev.yml exec -T api mockgen -source=./$(source)/$(filename) -destination=./mock/$(dest)/mock_$(filename) -package=$(package)
