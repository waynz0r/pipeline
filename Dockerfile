FROM golang:1.10-alpine

# warmup go build cache
ADD vendor /go/src/github.com/banzaicloud/pipeline/vendor
WORKDIR /go/src/github.com/banzaicloud/pipeline
RUN time go build ./vendor/... | true

ADD . /go/src/github.com/banzaicloud/pipeline
RUN time go build -o /pipeline main.go

FROM alpine:3.6
RUN apk add --no-cache ca-certificates
COPY --from=0 /pipeline /
COPY --from=0 /go/src/github.com/banzaicloud/pipeline/views /views/
ENTRYPOINT ["/pipeline"]
