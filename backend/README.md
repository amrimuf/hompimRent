# HompimRent API

A backend service for a rental system, built with Go and PostgreSQL. This service manages rental operations for items such as cars, cameras, bikes, and dresses.

## Prerequisites

-   **Docker:** Ensure Docker and Docker Compose are installed. You can follow the installation guide at [Docker Documentation](https://docs.docker.com/get-docker/).
-   **Go:** Version 1.18 or later is recommended. [Go Installation Guide](https://golang.org/doc/install).
  

## Setup

### 1. Clone the Repository

```bash
git clone https://github.com/amrimuf/hompimrent.git
cd hompimrent
```

### 2. Configure Environment Variables

Copy the example environment configuration file to `.env`:

```bash
cp env.example .env
```

Edit the `.env` file to add your PostgreSQL configuration:

```env
POSTGRES_USER=your_db_user
POSTGRES_PASSWORD=your_db_password
POSTGRES_DB=rental_system
```

### 3. Start the Services

Run Docker Compose to build and start the services:

```bash
docker-compose up --build
```

This command will start the PostgreSQL database and the backend application.

### 4. Apply Database Migrations

To apply database migrations, you can utilize the included migration service without needing to install golang-migrate manually. Just follow these steps:

```bash
docker-compose run migrate
```

This command will automatically run migrations defined in the `./database/migrations` directory.

### 5. Accessing the Application

-   **API Endpoints:** The application exposes various API endpoints for managing rentals. Check the API documentation for details.
-   **PostgreSQL Access:** Connect to PostgreSQL using the following command:

    ```bash
    docker exec -it your_postgres_container_name psql -U your_db_user
    ```

## Build and Run

### For Development:

To build and run the application in development mode, use the following command:

```bash
docker-compose -f docker-compose.yaml up --build
```

This command will build the Docker image using your development Dockerfile and start the application with the configurations defined in docker-compose.dev.yaml.

### For Production:

To build and run the application in production mode, use the following command:

```bash
docker-compose -f docker-compose.prod.yaml up --build
```

This command will build the Docker image using your production Dockerfile and start the application with the configurations defined in docker-compose.prod.yaml.

### Development

-   **Code:** The application code is in the `cmd/` directory. The main entry point is `cmd/main.go`.
-   **Testing:** Run tests using:

    ```bash
    go test ./...
    ```

### Troubleshooting

-   **Connection Issues:** Ensure Docker containers are running and accessible.
-   **Database Problems:** Verify that the `.env` file is correctly configured and that migrations are applied.

### License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

### Contact

For questions or issues, please contact [your.email@example.com](mailto:your.email@example.com).
