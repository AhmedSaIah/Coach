DOCKER_COMPOSE_FILE = docker-compose.yaml
DB_CONTAINER_NAME = taskmanager_db
DB_MIGRATE_CONTAINER_NAME = taskmanager_migrate
ENV_FILE = .env

GO = go
GOBUILD = $(GO) build
GORUN = $(GO) run
GOTEST = $(GO) test
GOCLEAN = $(GO) clean
GOFMT = $(GO) fmt
GOGET = $(GO) get
GOLINT = golangci-lint run

CMD_DIR = cmd

EXECUTABLE = taskmanager

.PHONY: all build run test clean fmt deps docker-up docker-down migrate dbshell

all: build

build:
	$(GOBUILD) -o $(EXECUTABLE) $(CMD_DIR)/main.go

run:
	$(GORUN) $(CMD_DIR)/main.go

test:
	$(GOTEST) ./...

clean:
	$(GOCLEAN)
	rm -f $(EXECUTABLE)

fmt:
	$(GOFMT) ./...

# Get dependencies
deps:
	$(GOGET) -v -t -d ./...

lint:
	$(GOLINT)

docker-up:
	docker-compose --env-file $(ENV_FILE) -f $(DOCKER_COMPOSE_FILE) up -d

docker-down:
	docker-compose --env-file $(ENV_FILE) -f $(DOCKER_COMPOSE_FILE) down

migrate:
	docker-compose --env-file $(ENV_FILE) -f $(DOCKER_COMPOSE_FILE) run --rm $(DB_MIGRATE_CONTAINER_NAME) migrate

# Open a shell in the DB container
dbshell:
	docker exec -it $(DB_CONTAINER_NAME) psql -U $$POSTGRES_USER -d $$POSTGRES_DB

rebuild: clean build
