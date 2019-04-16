FROM golang:1.9.7

MAINTAINER TripleZ "me@triplez.cn"

WORKDIR $GOPATH/src/a2os/behavior
ADD . $GOPATH/src/a2os/behavior/

# Build package
RUN go build -o behavior .

#EXPOSE 8080

ENTRYPOINT ["./behavior"]
