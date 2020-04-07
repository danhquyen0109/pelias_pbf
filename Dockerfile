FROM golang:1.13.7-alpine3.11 AS builder

WORKDIR "$GOPATH/src/github.com/danhquyen0109/pelias_pbf"

RUN apk update \
  && apk add git gcc musl-dev

COPY . "$GOPATH/src/github.com/danhquyen0109/pelias_pbf"

RUN go get && go build

FROM alpine:3.11.3

COPY --from=builder /go/src/github.com/danhquyen0109/pelias_pbf/pbf /bin/

ENTRYPOINT [ "pbf" ]
