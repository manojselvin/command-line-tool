# Command Line Tool

This is a command-line tool written in Go that consumes and displays TODO items from an API endpoint.

## Prerequisites

- Go (version 1.13 or later) installed on your machine
- Docker installed on your machine (optional, for running the application in a Docker container)

## Usage

### Local Development

1. Clone this repository to your local machine:

    ```bash
    git clone https://github.com/manojselvin/command-line-tool.git
    ```

2. Navigate to the project directory:

    ```bash
    cd command-line-tool
    ```

3. Build the Go application:

    ```bash
    go build -o command_line_tool .
    ```

4. Run the Go application:

    ```bash
    ./command_line_tool
    ```

### Docker

1. Clone this repository to your local machine (if you haven't already):

    ```bash
    git clone https://github.com/manojselvin/command-line-tool.git
    ```

2. Navigate to the project directory:

    ```bash
    cd command-line-tool
    ```

3. Build the Docker image:

    ```bash
    docker build -t command-line-tool .
    ```

4. Run the Docker container:

    ```bash
    docker run -p 8080:8080 command-line-tool
    ```

### Testing

To run tests for the application:

```bash
go test ./...
