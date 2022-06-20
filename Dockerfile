# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.18-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init

RUN go build -o /trading

##
## Deploy
##
FROM ubuntu:20.04

WORKDIR /

COPY --from=build /trading /trading
COPY --from=build /app/public/ /public/
COPY --from=build /app/docs/ /docs/
COPY --from=build /app/templates/ /templates/
ENV GIN_MODE release

EXPOSE 8080

ENTRYPOINT ["/trading"]