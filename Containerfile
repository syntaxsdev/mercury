FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build mercury
RUN go build -o main .

# Use a smaller image for running
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .

# Expose the web port
EXPOSE 8080

# Run mercury server
CMD ["./main"]
