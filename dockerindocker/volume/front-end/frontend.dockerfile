FROM alpine:latest

RUN mkdir /app

COPY frontApp /app
COPY cmd/web/templates /app/templates

WORKDIR /app

CMD [ "/app/frontApp"]
