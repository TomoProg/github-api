FROM golang:latest
MAINTAINER TomoProg <helloworld0306.xxx@gmail.com>
RUN apt-get update && apt-get install -y \
	vim-tiny
RUN go get "github.com/google/go-github/github"
WORKDIR /go/app
