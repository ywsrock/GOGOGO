FROM golang:latest

RUN mkdir /app
WORKDIR /app

ADD ./work /app

RUN go mod download
EXPOSE 8080

