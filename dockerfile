# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download && go build -o teen-wallet ./cmd/server

# Runtime stage
FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/teen-wallet /usr/local/bin/teen-wallet
COPY .env /app/.env
CMD ["teen-wallet", "--config", "/app/.env"]