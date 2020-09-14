FROM golang:1.15

ENV GOBIN /go/bin
ENV GO111MODULE=on
ENV GOPATH=

WORKDIR /go
ADD ./app /go

RUN go build -o main main.go
