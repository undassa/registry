FROM golang:1.11.1-alpine3.8

RUN apk add --no-cache ca-certificates

RUN apk add --no-cache iptables \
    linux-headers \
    gcc \
    musl-dev \
    iproute2

RUN set -ex \
	&& apk add --no-cache --virtual .build-deps \
    bash \
    git  \
    make \
	\
	&& rm -rf /*.patch

ENV GOPATH /go
ENV GOROOT /usr/local/go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

ADD . /go/src/github.com/lastbackend/registry
WORKDIR /go/src/github.com/lastbackend/registry

RUN make build && make install
RUN apk del --purge .build-deps

WORKDIR /go/bin
RUN rm -rf /go/pkg \
    && rm -rf /go/src \
    && rm -rf /var/cache/apk/*

EXPOSE 2954 2955