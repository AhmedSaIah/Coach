# TaskManagement

TaskManagement is a RESTful API for managing tasks, built with Go, PostgreSQL, Docker, and GORM.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Project Structure](#project-structure)
- [Setup](#setup)
- [Usage](#usage)
- [Makefile Commands](#makefile-commands)
- [API Endpoints](#api-endpoints)

## Prerequisites

- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/products/docker-desktop)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [golangci-lint](https://golangci-lint.run/usage/install/)

## Project Structure


## Setup

1. **Clone the repository**:
    ```sh
    git clone https://github.com/AhmedSaIah/TaskManager.git
    cd TaskManager
    ```

2. **Create the `.env` file**:
    ```sh
    cp .env.example .env
    ```

    Update the `.env` file with your database configuration:
    ```
    POSTGRES_DB=test
    POSTGRES_USER=test
    POSTGRES_PASSWORD=test
    DB_HOST=localhost
    DB_PORT=5432
    ```

3. **Start Docker containers**:
    ```sh
    make docker-up
    ```

4. **Run database migrations**:
    ```sh
    make migrate
    ```

5. **Build and run the application**:
    ```sh
    make build
    make run
    ```
    
## API Endpoints

- **Get all tasks**:
    ```sh
    GET /tasks
    ```

- **Get a task by ID**:
    ```sh
    GET /tasks/{id}
    ```

- **Create a new task**:
    ```sh
    POST /tasks
    ```

- **Update a task by ID**:
    ```sh
    PUT /tasks/{id}
    ```

- **Delete a task by ID**:
    ```sh
    DELETE /tasks/{id}
    ```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
