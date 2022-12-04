FROM golang:1.18.0-alpine3.15 as builder

RUN apk add --no-cache make musl-dev linux-headers gcc git jq bash

# build dl-challenger with local monorepo go modules
COPY ./middleware /app/middleware
COPY ./common /app/common
COPY ./lib /app/lib

WORKDIR /app/middleware/rollup-example/challenger

RUN make 


FROM alpine:3.15

RUN apk add netcat-openbsd

COPY --from=builder /app/middleware/rollup-example/challenger/bin/challenger /usr/local/bin

ENTRYPOINT ["challenger"]