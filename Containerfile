FROM golang:1.22.11-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./

# Download deps
RUN go mod download

# Copy source code
COPY . .

# Build mercury
RUN CGO_ENABLED=0 go build -o mercury main.go

# Use a smaller image for running
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/mercury .

# Expose the web port
EXPOSE 80

# Run mercury server
CMD ["./mercury"]
