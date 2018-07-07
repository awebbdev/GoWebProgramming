#To Build:  docker build -t awebbdev/chitchat .
#To Run:    docker run --rm -p 8080:8080 awebbdev/chitchat


FROM golang:latest

ENV SRC_LOC=github.com/awebbdev/GoWebProgramming/chitchat

ADD . /go/src/$SRC_LOC

RUN go install $SRC_LOC

ENTRYPOINT /go/bin/chitchat

EXPOSE 8080