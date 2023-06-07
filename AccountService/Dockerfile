FROM golang:1.20-alpine AS builder

RUN apk update && apk add --no-cache git ca-certificates

WORKDIR /app

RUN go get -d -v ./...
RUN go mod tidy

COPY . .

ARG TARGET_FILE=main.go

RUN go build -o /account-service ${TARGET_FILE}

FROM alpine

COPY .env /.env
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /account-service /

RUN adduser -D account-service
USER account-service

EXPOSE 50054

CMD ["/account-service"]
