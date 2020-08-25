FROM golang:stretch

ARG SSH_PRIVATE_KEY
ENV PKG_CONFIG_PATH /usr/lib/pkgconfig

RUN apt-get update \
 && apt-get install -y \
    curl \
    git \
    libcurl4-openssl-dev \
    ssh

COPY . /go/src/git.merostm.com/go-quote
WORKDIR /go/src/git.merostm.com/go-quote

RUN mkdir ~/.ssh && umask 0077 && \
    echo "${SSH_PRIVATE_KEY}" > ~/.ssh/id_rsa && \
    ssh-keyscan -p 7999 git.merostm.com >> ~/.ssh/known_hosts && \
    git config --global url."ssh://git@git.merostm.com:7999/goq/go-sse-quote-service.git".insteadOf "https://git.merostm.com/scm/goq/go-sse-quote-service" && \
    go get -d -v && \
    go build -o /go/bin/go_quote_service && \
    rm ~/.ssh/id_rsa

CMD ["/go/bin/go_quote_service"]