FROM golang:1.20-alpine AS builder

RUN apk update && apk add --no-cache git ca-certificates

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ARG TARGET_FILE=main.go

RUN go build -o /tweet-service ${TARGET_FILE}

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /tweet-service /

EXPOSE 50051

CMD ["/tweet-service"]