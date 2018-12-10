FROM waynz0r/pipeline-build-base:latest AS depcache

RUN date

FROM golang:1.11-alpine as build

RUN date

RUN apk add --update --no-cache bash ca-certificates curl git make
RUN go get -d github.com/kubernetes-sigs/aws-iam-authenticator/cmd/aws-iam-authenticator
RUN cd $GOPATH/src/github.com/kubernetes-sigs/aws-iam-authenticator && \
    git checkout 981ecbe && \
    go install ./cmd/aws-iam-authenticator

RUN date

COPY --from=depcache /go/pkg /go/pkg

RUN date

RUN mkdir -p /go/src/github.com/banzaicloud/pipeline
ADD Gopkg.* Makefile /go/src/github.com/banzaicloud/pipeline/
WORKDIR /go/src/github.com/banzaicloud/pipeline
RUN make vendor

RUN date

ADD . /go/src/github.com/banzaicloud/pipeline

RUN date

RUN BUILD_DIR=/ make build

RUN date

FROM alpine:3.7

RUN date

RUN apk add --update --no-cache tzdata

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/bin/aws-iam-authenticator /usr/bin/
COPY --from=build /go/src/github.com/banzaicloud/pipeline/views /views/
COPY --from=build /pipeline /

RUN date

ENTRYPOINT ["/pipeline"]
