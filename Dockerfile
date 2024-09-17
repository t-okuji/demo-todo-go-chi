FROM golang:1.23.0-alpine
WORKDIR /app
RUN apk add --no-cache ca-certificates && update-ca-certificates
RUN go install github.com/air-verse/air@latest