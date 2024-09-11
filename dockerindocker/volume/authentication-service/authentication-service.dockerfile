FROM alpine:latest

RUN mkdir -p /app

COPY authApp /app

CMD [ "/app/authApp"]