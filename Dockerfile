# syntax=docker/dockerfile:1

## Build

FROM golang:1.18-alpine AS build

WORKDIR /app

COPY . /app/

RUN go mod download
WORKDIR /app/cmd

RUN go build -o /github-interaction-v2


## Deploy
FROM alpine:latest

WORKDIR /

COPY --from=build /github-interaction-v2 /github-interaction-v2
EXPOSE 8080

ENTRYPOINT ["/github-interaction-v2"]