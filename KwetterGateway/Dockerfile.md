FROM golang:1.20-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ARG TARGET_FILE=main.go

RUN go build -o /kwetter-gateway ${TARGET_FILE}

FROM scratch

COPY --from=builder /kwetter-gateway /

EXPOSE 8080

CMD ["/kwetter-gateway"]