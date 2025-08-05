## Project Overview

This project is a Go-based backend service that provides a comprehensive CRUD API for managing users, products, orders, and inventory. It is built using the Echo framework, a high-performance, extensible, and minimalist web framework for Go. For its database, the application uses MongoDB, a popular NoSQL database, and interacts with it via the official MongoDB Go driver.

The application is containerized using Docker, with a `Dockerfile` for building the Go application and a `docker-compose.yml` file for orchestrating the application and a MongoDB instance. This setup ensures a consistent and reproducible development and deployment environment.

The project follows a standard Go project layout, with a clear separation of concerns:

- **`cmd/api`**: Contains the main application entry point.
- **`internal/handlers`**: Defines the HTTP handlers and routing logic.
- **`internal/models`**: Specifies the data structures for the application.
- **`pkg/database`**: Manages the database connection.

## Building and Running

To build and run the project, you will need to have Docker and Docker Compose installed on your system. Once you have them, you can use the following command to start the application:

```bash
docker-compose up --build
```

This command will build the Go application, start the MongoDB container, and make the API available at `http://localhost:8080`.

## Development Conventions

- **Framework**: The project uses the Echo framework for routing and handling HTTP requests.
- **Database**: MongoDB is the database of choice, and the official Go driver is used for all database interactions.
- **Dependency Management**: Go Modules are used for managing project dependencies.
- **Containerization**: The application is containerized with Docker, and Docker Compose is used for orchestration.
- **Code Style**: The project follows standard Go coding conventions.
