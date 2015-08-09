FROM golang

WORKDIR /gopath/app
ENV GOPATH /gopath/app
#ADD . /gopath/app/

RUN go get github.com/shaalx/news

EXPOSE 80
CMD ["/gopath/app/bin/news"]
