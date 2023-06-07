FROM golang:1.20-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG TARGET_FILE=main.go

RUN CGO_ENABLED=0 go build -o /gateway ${TARGET_FILE}

FROM alpine

COPY --from=builder /gateway /

EXPOSE 8080
EXPOSE 50055

CMD ["/gateway"]
