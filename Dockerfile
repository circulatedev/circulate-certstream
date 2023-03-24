# Use the official Golang image as the base image
FROM golang:1.19-alpine as builder

# Set the working directory
WORKDIR /app

# Copy the Go modules files
COPY go.mod .
COPY go.sum .

# Download the dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o circulate-certstream

# Create a new lightweight alpine image
FROM alpine:3.14

RUN addgroup -g 1000 gogroup && \
    adduser -u 1000 -G gogroup -s /bin/sh -D gouser

# Copy the binary from the builder stage
COPY --from=builder /app/circulate-certstream /app/circulate-certstream

# Set the working directory
WORKDIR /app

# Expose the server port
EXPOSE 8080

RUN chown gouser:gogroup /app/circulate-certstream && \
    chmod 750 /app/circulate-certstream

USER 1000
    
# Run the application
CMD ["/app/circulate-certstream"]
