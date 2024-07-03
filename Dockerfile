FROM golang:latest

WORKDIR /usr/src/app
# WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY . .
RUN go mod tidy