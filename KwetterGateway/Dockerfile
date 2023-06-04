FROM golang:1.20-alpine AS builder

RUN apk update && apk add --no-cache git ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG TARGET_FILE=main.go

RUN go build -o /gateway ${TARGET_FILE}
COPY .env /app/.env

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /gateway /
COPY --from=builder /app/.env .

EXPOSE 8080
EXPOSE 50055

CMD ["/gateway"]
