# Stage 1: Build the Go binary
FROM golang:1.22 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main cmd/gate/main.go

# Stage 2: Create the final lightweight image with Alpine
FROM alpine:3.18

# Install CA certificates
RUN apk add --no-cache ca-certificates

# Copy the binary and environment file from the builder stage
COPY --from=builder /main /main
COPY --from=builder /app/.env /.env

# Command to run the executable
ENTRYPOINT ["/main"]