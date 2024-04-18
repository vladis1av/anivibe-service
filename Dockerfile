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

FROM nginx:alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates

COPY --from=builder /usr/share/zoneinfo/America/New_York /usr/share/zoneinfo/America/New_York
ENV TZ America/New_York

WORKDIR /app
COPY --from=builder /app/app .

# Копируем конфигурацию Nginx
COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 80

# Запускаем приложение и Nginx
CMD ./app & nginx -g 'daemon off;'