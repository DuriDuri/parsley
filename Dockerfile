FROM golang:1.7

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

CMD ["go-wrapper", "run"]

COPY . /go/src/app
RUN go-wrapper install
EXPOSE 8080
