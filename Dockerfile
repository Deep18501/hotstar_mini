# syntax=docker/dockerfile:1

# === Stage 1: Builder ===
# Use the official Golang image for building the application.
FROM golang:1.25.5-alpine3.23 AS builder

# Set the working directory inside the container.
WORKDIR /app

# Copy go.mod and go.sum files first to leverage Docker's build cache.
COPY go.mod go.sum ./

# Download all the dependencies.
RUN go mod download

# Copy the rest of the application source code.
COPY . .

# Build the Go application binary with default name.
# CGO_ENABLED=0 disables CGo, making the binary static and portable.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /app/hotstar_mini . && \
    ls -lah /app/hotstar_mini


# === Stage 2: Runtime ===
# Use Alpine Linux for a minimal but functional container.
FROM alpine:3.23

# Install dependencies and Go
RUN apk add --no-cache bash curl wget && \
    wget -q https://go.dev/dl/go1.25.5.linux-amd64.tar.gz -O go.tar.gz && \
    tar -C /usr/local -xzf go.tar.gz && \
    rm go.tar.gz

# Copy entire app directory from builder
COPY --from=builder /app /app

# Set the working directory
WORKDIR /app

# Setup PATH for Go and tools installed via exec
ENV PATH="/usr/local/go/bin:/root/go/bin:${PATH}"

# Expose the port your application listens on
EXPOSE 8080

# Keep container running for manual execution via docker exec
CMD ["tail", "-f", "/dev/null"]
