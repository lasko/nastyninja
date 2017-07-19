FROM golang:1.4.2-wheezy

WORKDIR /go/src/lambda
ADD . /go/src/lambda

RUN go get -v -d
RUN go install -v

EXPOSE 5000
CMD []
ENTRYPOINT ["/go/bin/lambda"]
