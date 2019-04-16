FROM golang:1.9.7

MAINTAINER TripleZ "me@triplez.cn"

# Solution for Chinese special network enviornment
# RUN mkdir -p $GOPATH/src/golang.org/x/ && \
#     git clone https://github.com/golang/sys.git $GOPATH/src/golang.org/x/sys/
RUN mkdir -p $GOPATH/src/golang.org/x/ && \
    git clone https://gitee.com/Triple-Z/golang-sys.git $GOPATH/src/golang.org/x/sys/

# Add dependencies
# RUN go get -u github.com/kardianos/govendor && go get github.com/pilu/fresh && govendor sync

# Add fresh package
RUN go get github.com/pilu/fresh

# Add behavior code
WORKDIR $GOPATH/src/a2os/behavior
ADD . $GOPATH/src/a2os/behavior/

# Build package
RUN go build -o behavior-dev .

#EXPOSE 8080

ENTRYPOINT ["fresh"]
#ENTRYPOINT ["./behavior-dev"]
