FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o zipweather ./cmd/main.go

# Final minimal image
FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/zipweather .
COPY .env .env

# Default entrypoint
ENTRYPOINT ["./zipweather"]