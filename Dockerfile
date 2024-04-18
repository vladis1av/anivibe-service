# Этап сборки Go приложения
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
FROM alpine:latest

# Устанавливаем NGINX Unit
RUN apk add --no-cache curl gnupg \
    && curl -fsSL https://unit.nginx.com/keys/nginx-unit.key | gpg --dearmor | dd of=/etc/apk/keys/nginx-unit.rsa.pub \
    && echo "http://unit.nginx.com/deb/alpine/latest/main" >> /etc/apk/repositories \
    && apk add --no-cache nginx-unit

RUN apk update --no-cache && apk add --no-cache ca-certificates

COPY --from=builder /usr/share/zoneinfo/America/New_York /usr/share/zoneinfo/America/New_York
ENV TZ America/New_York

WORKDIR /app
COPY --from=builder /app/app .

# Копируем конфигурацию NGINX Unit
COPY unit.json /docker-entrypoint.d/

# Открываем порт для NGINX Unit
EXPOSE 8080

# Устанавливаем переменные окружения
ENV HTTP_ADDR=:8080
ENV READ_TIMEOUT=5s
ENV WRITE_TIMEOUT=10s
ENV IDLE_TIMEOUT=60s
ENV MAX_HEADER_BYTES=1048576
ENV PROXY_HTTP_ADDR=:8081
ENV PROXY_READ_TIMEOUT=30s
ENV PROXY_WRITE_TIMEOUT=30s
ENV PROXY_IDLE_TIMEOUT=60s
ENV PROXY_MAX_HEADER_BYTES=1048576

# Запускаем приложение и NGINX Unit
CMD ./app & unitd --no-daemon --control unix:/var/run/control.unit.sock