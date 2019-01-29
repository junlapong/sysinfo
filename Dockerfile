FROM golang:1.11

RUN apt-get update && \
    go get github.com/buger/goterm

COPY sysinfo.go /go/src
