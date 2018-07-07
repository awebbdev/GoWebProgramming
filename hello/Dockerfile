#To Build:  docker build -t awebbdev/hello .
#To Run:    docker run --rm -p 8080:8080 awebbdev/hello


FROM golang:latest

ENV SRC_LOC=github.com/awebbdev/GoWebProgramming/hello

ADD . /go/src/$SRC_LOC

RUN go install $SRC_LOC

ENTRYPOINT /go/bin/hello

EXPOSE 8080