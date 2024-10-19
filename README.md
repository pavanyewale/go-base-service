# Go Base Service

This is a simple service designed to demonstrate how to structure code in Go.

<div align="start">
<img src="golang.png" alt="Go Logo" width="200"/>
</div>

## Prerequisites

- Go version 1.22.4 or above installed
- buf is installed 
   ``` 
   brew install bufbuild/buf/buf
   ```

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

## Development Notes

- **Configuration**: The configuration settings can be managed in `internal/configs/configs.go`.
- **Logging**: Ensure proper logging is implemented for easier debugging and monitoring.
- **Error Handling**: Consistent error handling throughout the codebase.

## Developers and contact
Pavan Yewale
email: pavanyewale1996@gmail.com

