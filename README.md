================================================================================================================================================================
To install go locally::::::::========>>>>>>>>>>
wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz
mkdir -p .go_runtime
tar -C .go_runtime -xzf go1.22.0.linux-amd64.tar.gz
rm go1.22.0.linux-amd64.tar.gz
export GOROOT=$(pwd)/.go_runtime/go
export PATH=$GOROOT/bin:$PATH
export GOPATH=$(pwd)/.go_path

****************************************************************************************************************************************************************





# Go Product API

A REST API for User Authentication and Product CRUD operations, built with Go, Gin, Gorm, and SQLite.

## Prerequisites

- Go 1.21 or higher

## Setup

1.  **Configure Local Go Environment:**

    Since a system-wide Go installation was not found, a local version has been installed in `.go_runtime`. Run the following command to set up your environment:

    ```bash
    source ./setup_env.sh
    ```

2.  **Initialize Module & Install Dependencies:**

    Once the environment is sourced, run:

    ```bash
    go mod tidy
    ```

3.  **Configuration:**

    A `.env` file is provided with default values:
    ```env
    PORT=:8080
    DB_URL=test.db
    JWT_SECRET=supersecretkey
    ```

## Running the Application

```bash
go run cmd/api/main.go
```

## API Endpoints

### Authentication

-   **Register:** `POST /api/auth/register`
    ```json
    {
      "email": "user@example.com",
      "password": "password123"
    }
    ```
-   **Login:** `POST /api/auth/login`
    ```json
    {
      "email": "user@example.com",
      "password": "password123"
    }
    ```
    *Returns a JWT token.*

### Products

-   **Get All Products:** `GET /api/products/`
-   **Get Product by ID:** `GET /api/products/:id`

**Protected Routes (Requires `Authorization: Bearer <token>` header):**

-   **Create Product:** `POST /api/products/`
    ```json
    {
      "name": "Laptop",
      "description": "Powerful laptop",
      "price": 999.99,
      "user_id": 1
    }
    ```
-   **Update Product:** `PUT /api/products/:id`
-   **Delete Product:** `DELETE /api/products/:id`

## Production Features

-   **Graceful Shutdown:** The server handles `SIGINT` and `SIGTERM` signals for clean shutdown.
-   **Structured Logging:** JSON logging via `slog` for better observability.
-   **Rate Limiting:** IP-based rate limiting to prevent abuse.
-   **CORS:** Configured for cross-origin requests.
-   **Docker:** Multi-stage build for small, secure images.
-   **Makefile:** Shortcuts for common tasks.

## Commands

-   `make run`: Run locally.
-   `make build`: Build binary.
-   `make test`: Run tests.
-   `make docker-build`: Build Docker image.
-   `make docker-run`: Run Docker container.




