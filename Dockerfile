FROM golang

WORKDIR /gopath/app
ENV GOPATH /gopath/app
ADD . /gopath/app/

RUN go get github.com/toukii/news

EXPOSE 8080:80
CMD ["/gopath/app/bin/news"]
