FROM golang:1.18.3-alpine3.15 as build

WORKDIR /app
COPY . .
RUN go build -o main
RUN apk add busybox-extras

