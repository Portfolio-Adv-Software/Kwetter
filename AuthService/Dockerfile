FROM golang:1.20-alpine AS builder

RUN apk update && apk add --no-cache git ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . /app

ARG TARGET_FILE=main.go

RUN go build -o /auth-service ${TARGET_FILE}

FROM scratch

COPY .env /.env
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /auth-service /

RUN adduser -D auth-service
USER auth-service

EXPOSE 50053

CMD ["/auth-service"]