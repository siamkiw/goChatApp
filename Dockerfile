#FROM golang:1.16
#
#RUN mkdir /app
#
#ADD . /app

FROM golang:1.16

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]