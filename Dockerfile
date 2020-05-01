FROM golang:alpine

ADD . $GOPATH/src/pando
WORKDIR $GOPATH/src/pando

RUN apk add --update --no-cache git gcc libc-dev tzdata;
#  tzdata wget gcc libc-dev make openssl py-pip;

ENV TZ=${PANDO_TIMEZONE}

CMD go get \
    && go build -v -o $GOPATH/bin/pando \
    && if [ $PANDO_NODE == "server" ]; then \
        pando run server ; \ 
    elif [ $PANDO_NODE == "worker" ]; then \
        pando run worker ; \
    fi

EXPOSE ${PANDO_PORT}