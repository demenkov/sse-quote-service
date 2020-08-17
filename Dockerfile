FROM golang:stretch

ENV PKG_CONFIG_PATH /usr/lib/pkgconfig

RUN apt-get update \
 && apt-get install -y \
    curl \
    git \
    libcurl4-openssl-dev

COPY . /go/src/git.merostm.com/go-quote
WORKDIR /go/src/git.merostm.com/go-quote

RUN go get -d -v \
 && go build -o /go/bin/go_quote_service

CMD ["/go/bin/go_quote_service"]