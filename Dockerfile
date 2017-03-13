FROM golang:1.8.0-alpine

COPY . /go/src/github.com/uudashr/fibweb

WORKDIR /go/src/github.com/uudashr/fibweb

RUN apk update && apk upgrade && \
    apk --no-cache --update add git && \
    go get -d -v ./... && go install -v ./... && \
    apk del git && rm -rf /var/cache/apk/*

EXPOSE 8080

CMD fibweb -fibgo-addr http://$FIBGO_ADDR
