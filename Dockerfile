FROM golang:1.20.1

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod tidy
