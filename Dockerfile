# Base
FROM golang:1.23.0-alpine AS base
WORKDIR /app

# Development
FROM base AS dev
RUN go install github.com/air-verse/air@latest

# Builder for prod
FROM base AS builder
COPY ./app/ /app/
RUN go build -o main

# Production
FROM alpine:latest AS prod
WORKDIR /go/src/app
COPY --from=builder /app/main ./main
ENTRYPOINT [ "./main" ]