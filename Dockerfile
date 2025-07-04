# -------- Builder Stage --------
FROM golang:1.24.1 AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (for dependency caching)
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the full source code
COPY . .

# Build the Go app
RUN go build -o todo-app

# -------- Final Stage --------
FROM debian:bookworm

# Set working directory in the final container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/todo-app .

# Set entry point
ENTRYPOINT ["./todo-app"]

# Expose port (adjust if needed)
EXPOSE 8080
