# Start from the latest golang base image
FROM golang:1.19.3 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./
COPY go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY ./ ./

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/main

# Start a new stage from scratch
FROM gcr.io/distroless/static-debian11:nonroot

# Set the Current Working Directory inside the container
WORKDIR /

# Copy the Pre-built binary file from the previous stage
COPY --from=builder --chown=nonroot:nonroot /app/main /main

COPY --from=builder --chown=nonroot:nonroot /app/.env /.env

# Expose port 8080 to the outside world
EXPOSE 8080

# Create nonroot user
USER nonroot:nonroot

# Command to run the executable
ENTRYPOINT ["/main"]