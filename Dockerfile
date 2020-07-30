FROM golang:1.14-alpine3.11 as base

FROM base as build
RUN apk add --update make

ENV GO111MODULE=on
COPY . /root/src/go/github.com/ninth-wave/nwp-load-tester
WORKDIR /root/src/go/github.com/ninth-wave/nwp-load-tester
RUN make build


FROM alpine:3.11 as runtime
RUN apk update && apk upgrade && apk add bash && apk add curl && apk add perl

COPY --from=build /root/src/go/github.com/ninth-wave/nwp-load-tester/bin/load-tester /usr/local/bin

WORKDIR /usr/local/bin

ENTRYPOINT ["/usr/local/bin/load-tester"]