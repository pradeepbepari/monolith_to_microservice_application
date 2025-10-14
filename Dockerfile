# Use official Golang image based on Alpine Linux for minimal footprint
FROM golang:1.24-alpine AS builder

# Set working directory inside the container
WORKDIR /app

# Install Air for live reloading during development
RUN go install github.com/air-verse/air@v1.61.7

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download 

# Copy the rest of the application source code
COPY . .

# Start the application using Air with the specified config
CMD ["air", "-c", ".air.toml"]
