FROM golang:1.20-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG TARGET_FILE=main.go

RUN CGO_ENABLED=0 go build -o /trend-service ${TARGET_FILE}

FROM alpine

COPY .env /.env
COPY --from=builder /trend-service /

RUN adduser -D trend-service
USER trend-service

EXPOSE 50052

CMD ["/trend-service"]