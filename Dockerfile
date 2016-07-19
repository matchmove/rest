# Base the docker container off of the official golang image
FROM golang:latest

# Install dependencies
RUN go get github.com/gorilla/mux
RUN go get gopkg.in/yaml.v2

# Copy the local package files to the container’s workspace.
ADD . /go/src/bitbucket.org/matchmove/rest

# Install api binary globally within container
RUN go install bitbucket.org/matchmove/rest

# Run test cases
RUN go test bitbucket.org/matchmove/rest -v