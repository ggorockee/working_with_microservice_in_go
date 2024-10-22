## base go image
FROM golang:1.23-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o campingApp ./cmd/api

RUN chmod +x /app/campingApp

## build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/campingApp /app
COPY ./cmd/api/docs /app
RUN ls /app
RUN ls /app/docs

CMD ["/app/campingApp"]