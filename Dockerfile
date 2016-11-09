FROM golang:1.7
MAINTAINER "Neville Tummon"
EXPOSE 5000
RUN mkdir -p /go/src/titanlogger
COPY . /go/src/titanlogger
WORKDIR /go/src/titanlogger
RUN go get && go build
CMD titanlogger




