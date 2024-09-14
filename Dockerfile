# Build the Go application
FROM golang:1.23-alpine AS build

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Tidy up dependencies
RUN go mod tidy

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN go build -o main .

# Start a new minimal container
FROM alpine:latest
WORKDIR /root/

# Install bash
RUN apk add --no-cache bash

# Copy built app and wait-for-it script
COPY --from=build /app/main .
COPY wait-for-it.sh .
RUN chmod +x wait-for-it.sh

EXPOSE 8080
CMD ["./wait-for-it.sh", "db:5432", "--", "./main"]
