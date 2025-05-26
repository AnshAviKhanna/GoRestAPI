# Stage 1: Build the Go binary using official Golang image
FROM golang:1.20 as builder

WORKDIR /app

# Copy all necessary files
COPY . .

# Build the Go binary for Linux explicitly
RUN GOOS=linux GOARCH=amd64 go build -o server

# Stage 2: Use Alpine to run the binary
FROM alpine:latest

WORKDIR /root/

# Install libc if needed (required for Go binaries in Alpine)
RUN apk add --no-cache libc6-compat

# Copy the binary from builder
COPY --from=builder /app/server .

EXPOSE 8080

# Run the binary
CMD ["./server"]
