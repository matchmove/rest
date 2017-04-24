FROM golang:1.8

ENV APPDIR $GOPATH/src/github.com/matchmove/rest

ADD . $APPDIR
WORKDIR $APPDIR

RUN go get ./...
RUN go get github.com/gorilla/handlers
