FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o app ./cmd/app

# Финальный этап
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]