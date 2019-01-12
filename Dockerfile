FROM golang:1.11

RUN apt-get update; \
    cd /go/src && go get github.com/buger/goterm

COPY sysinfo.go /go/src
