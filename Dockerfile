# Start from the official golang base image
FROM golang:latest AS build

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Run tests
RUN go test ./...

# Build the Go application if tests succeed
RUN GOOS=linux GOARCH=amd64 go build -o command_line_tool .

# Start a new stage from scratch
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /app

# Copy the binary from the build stage to the final stage
COPY --from=build /app/command_line_tool .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./command_line_tool"]
