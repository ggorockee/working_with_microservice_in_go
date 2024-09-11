FROM docker:27.2.1-dind

RUN apk update &&\
    apk add curl make

RUN mkdir -p /temp &&\
    cd /temp &&  curl -LO https://go.dev/dl/go1.23.0.linux-arm64.tar.gz &&\
    rm -rf /usr/local/go && tar -C /usr/local -xzf go1.23.0.linux-arm64.tar.gz &&\
    rm -rf /temp

ENV PATH=$PATH:/usr/local/go/bin
ENV DOCKER_TLS_CERTDIR=/certs