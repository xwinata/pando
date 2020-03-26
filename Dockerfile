FROM golang:alpine

ADD . $GOPATH/src/pando
WORKDIR $GOPATH/src/pando

RUN apk add --update --no-cache git gcc libc-dev tzdata;
#  tzdata wget gcc libc-dev make openssl py-pip;

ENV TZ=${PANDO_TIMEZONE}

CMD go get \
    && go build -v -o $GOPATH/bin/pando \
    && pando run server;
EXPOSE ${PANDO_PORT}