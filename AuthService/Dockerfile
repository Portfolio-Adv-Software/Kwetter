FROM golang:1.20-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG TARGET_FILE=main.go

RUN go build -o /auth-service ${TARGET_FILE}

FROM golang:1.20-alpine

COPY .env /.env
COPY --from=builder /auth-service /

RUN adduser -D auth-service
USER auth-service

EXPOSE 50053

CMD ["/auth-service"]