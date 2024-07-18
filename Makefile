# Variables
DOCKER_COMPOSE_FILE = docker-compose.yaml
DB_CONTAINER_NAME = taskmanager_db
DB_MIGRATE_CONTAINER_NAME = taskmanager_migrate
ENV_FILE = .env

# Go commands
GO = go
GOBUILD = $(GO) build
GORUN = $(GO) run
GOTEST = $(GO) test
GOCLEAN = $(GO) clean
GOFMT = $(GO) fmt
GOGET = $(GO) get
GOLINT = golangci-lint run

# Directories
CMD_DIR = cmd

# Executable name
EXECUTABLE = taskmanager

.PHONY: all build run test clean fmt deps docker-up docker-down migrate dbshell

all: build

# Build the Go application
build:
	$(GOBUILD) -o $(EXECUTABLE) $(CMD_DIR)/main.go

# Run the Go application
run:
	$(GORUN) $(CMD_DIR)/main.go

# Run tests
test:
	$(GOTEST) ./...

# Clean the build
clean:
	$(GOCLEAN)
	rm -f $(EXECUTABLE)

# Format the Go code
fmt:
	$(GOFMT) ./...

# Get dependencies
deps:
	$(GOGET) -v -t -d ./...

# Linter
lint:
	$(GOLINT)

# Bring up Docker containers
docker-up:
	docker-compose --env-file $(ENV_FILE) -f $(DOCKER_COMPOSE_FILE) up -d

# Bring down Docker containers
docker-down:
	docker-compose --env-file $(ENV_FILE) -f $(DOCKER_COMPOSE_FILE) down

# Migrate the database
migrate:
	docker-compose --env-file $(ENV_FILE) -f $(DOCKER_COMPOSE_FILE) run --rm $(DB_MIGRATE_CONTAINER_NAME) migrate

# Open a shell in the DB container
dbshell:
	docker exec -it $(DB_CONTAINER_NAME) psql -U $$POSTGRES_USER -d $$POSTGRES_DB

# Rebuild the application
rebuild: clean build
