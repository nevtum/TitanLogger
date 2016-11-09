FROM golang:1.7
MAINTAINER "Neville Tummon"
COPY . /go/src/titanlogger
WORKDIR /go/src/titanlogger
RUN go get && go build
EXPOSE 5000
CMD titanlogger




