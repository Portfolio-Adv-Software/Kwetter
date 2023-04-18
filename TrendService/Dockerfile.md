# Build stage
FROM golang:1.20-alpine AS build
WORKDIR /app
COPY . .
RUN go build -o /go/bin/app

# Runtime stage
FROM alpine:latest
WORKDIR /app
COPY --from=build /go/bin/app .
CMD ["./app"]

# Expose the application port
EXPOSE 50051

# Expose the application port
EXPOSE 50052