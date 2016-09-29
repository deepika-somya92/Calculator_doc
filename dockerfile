FROM golang:latest

RUN go get github.com/fedolin/calculator

WORKDIR /go/src/github.com/fedolin/calculator

EXPOSE 9090

CMD ["calculator"]
