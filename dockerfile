# Use the official Golang image to build the app
FROM golang:1.20 as builder

WORKDIR /app

# Copy only go.mod (no go.sum needed)
COPY go.mod ./
RUN go mod tidy

# Copy the rest of the source code
COPY . .

# Build the Go binary
RUN go build -o server

# Use a minimal base image to run the app
FROM alpine:latest

WORKDIR /root/

# Copy the compiled binary from the builder
COPY --from=builder /app/server .

# Expose port 8080
EXPOSE 8080

# Run the app
CMD ["./server"]
