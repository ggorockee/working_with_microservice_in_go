FROM docker:27.2.1-dind-alpine3.20

# 필요한 라이브러리 설치
RUN apk add make curl


# go binary 다운로드
RUN mkdir -p /temp &&\
    cd /temp && curl -LO https://go.dev/dl/go1.23.0.linux-amd64.tar.gz

# Remove any previous Go installation by deleting the /usr/local/go folder (if it exists), \
#    then extract the archive you just downloaded into /usr/local, creating a fresh Go tree in /usr/local/go:
RUN  rm -rf /usr/local/go && tar -C /usr/local -xzf /temp/go1.23.0.linux-amd64.tar.gz && \
     rm -rf /temp

ENV PATH="$PATH:/usr/local/go/bin"

RUN go version