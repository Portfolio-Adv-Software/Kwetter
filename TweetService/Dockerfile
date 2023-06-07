FROM golang:1.20-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG TARGET_FILE=main.go

RUN go build -o /tweet-service ${TARGET_FILE}

FROM golang:1.20-alpine

COPY .env /.env
COPY --from=builder /tweet-service /

RUN adduser -D tweet-service
USER tweet-service

EXPOSE 50051

CMD ["/tweet-service"]