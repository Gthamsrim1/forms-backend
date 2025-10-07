# ---------- Build stage ----------
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go binary
RUN go build -o forms-server ./main.go


# ---------- Runtime stage ----------
FROM alpine:latest

WORKDIR /app

# Copy compiled binary
COPY --from=builder /app/forms-server .

# Copy essential runtime files (migrations, .env, etc.)
COPY migrations ./migrations
COPY .env .env

# Expose your app port
EXPOSE 8000

# Run the server
CMD ["./forms-server"]
