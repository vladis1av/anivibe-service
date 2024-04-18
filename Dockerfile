FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o /app/app ./cmd/app

# Финальный этап
FROM caddy:2.6.2-alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates

COPY --from=builder /usr/share/zoneinfo/America/New_York /usr/share/zoneinfo/America/New_York
ENV TZ America/New_York

WORKDIR /app
COPY --from=builder /app/app .

# Копируем конфигурацию Caddy
COPY Caddyfile /etc/caddy/Caddyfile

# Открываем порт 80 для Caddy
EXPOSE 80
# Устанавливаем переменные окружения
ENV HTTP_ADDR=:8081
ENV READ_TIMEOUT=5s
ENV WRITE_TIMEOUT=10s
ENV IDLE_TIMEOUT=60s
ENV MAX_HEADER_BYTES=1048576
ENV PROXY_HTTP_ADDR=:8082
ENV PROXY_READ_TIMEOUT=30s
ENV PROXY_WRITE_TIMEOUT=30s
ENV PROXY_IDLE_TIMEOUT=60s
ENV PROXY_MAX_HEADER_BYTES=1048576

# Запускаем два экземпляра приложения и Caddy
CMD ./app --http-addr=$HTTP_ADDR & ./app --http-addr=$PROXY_HTTP_ADDR & caddy run --config /etc/caddy/Caddyfile