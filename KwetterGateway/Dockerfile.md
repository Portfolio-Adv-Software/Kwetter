# Build stage
FROM golang:1.20-alpine AS build

# Copy the go.mod and go.sum files from the root directory of your monorepo into the container
WORKDIR /app
COPY Kwetter/go.mod Kwetter/go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the Go app
RUN go build -o /go/bin/app

# Runtime stage
FROM alpine:latest
WORKDIR /app
COPY --from=build /go/bin/app .
CMD ["./app"]