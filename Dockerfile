FROM golang:1.18.0-alpine3.15 as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git jq bash

COPY ./da-challenger /go/da-challenger
COPY ./bss-core /go/bss-core
COPY ./l2geth /go/l2geth

WORKDIR /go/da-challenger
RUN make

FROM alpine:3.15

RUN apk add --no-cache ca-certificates jq curl
COPY --from=builder /go/da-challenger/da-challenger /usr/local/bin/

WORKDIR /usr/local/bin

ENTRYPOINT ["da-challenger"]
