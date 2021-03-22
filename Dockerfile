FROM golang:latest
#RUN apk update && apk add git
RUN mkdir /go/src/app
WORKDIR /go/src/app
ADD . /go/src/app

RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/jinzhu/gorm