# Go Base Service

This is a simple service designed to demonstrate how to structure code in Go.

<div align="start">
<img src="golang.png" alt="Go Logo" width="200"/>
</div>

## Prerequisites

- Go version 1.22.4 or above installed

## How to Run

### Local Development

1. **Setup**: Install dependencies and prepare the environment.
    ```bash
    make setup
    ```

2. **Run the Service**: Start the service locally.
    ```bash
    make run
    ```

### Using Docker

1. **Build and Run with Docker**: Create a Docker image and run the service inside a container.
    ```bash
    make docker
    make docker-run
    ```

## Project Structure

Here's a basic overview of the project structure:
```bash
├── controller/ # Private application and library code
│ ├── grpc/ # GRPC controller
│ ├── http/ # HTTP controller
│ └── controller.go # is initialisation of controller is here
├── service/ # business logic
│ ├── service1/ # service 1 logic is here
│ ├── ping/ # ping service logic is here
│ ├── ...
│ └── service.go # is initialisation of service business logic factory is here
├── repository/ # data store
│ ├── mongo/ # mongo repository
│ ├── mssql/ # mssql repository
│ ├── ...
│ └── repository.go # switch of the repository is here
├── configs/ # Configuration files
│ └── configs.go # Configuration management
├── docs/ # Documentation files
│ └── swagger/ # Swagger files for API documentation
├── Makefile # Makefile for running tasks
└── README.md # Project overview (this file)
└── Dockerfile # for dockerise the application
└── main.go # start of the application
```

## Development Notes

- **Configuration**: The configuration settings can be managed in `configs/configs.go`.
- **Logging**: Ensure proper logging is implemented for easier debugging and monitoring.
- **Error Handling**: Consistent error handling throughout the codebase.

## Developers and contact
Pavan Yewale
email: pavanyewale1996@gmail.com

