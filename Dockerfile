FROM golang:alpine AS gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o /app/app ./cmd/app

FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates

COPY --from=gobuilder /usr/share/zoneinfo/Europe/Moscow /usr/share/zoneinfo/Europe/Moscow
ENV TZ Europe/Moscow

WORKDIR /app
COPY --from=gobuilder /app/app .

EXPOSE 80

CMD ["./app"]