#######################
# Build stage         #
#######################
FROM golang:1.22-alpine AS builder

WORKDIR /app


ENV GOTOOLCHAIN=auto

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o teen-wallet-api ./cmd/server


# Runtime stage       #

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/teen-wallet-api .
EXPOSE 8080
CMD ["./teen-wallet-api"]
