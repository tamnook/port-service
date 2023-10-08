FROM golang:1.21

COPY . /go/src/app

WORKDIR /go/src/app/cmd/port-service

RUN go build -o app main.go

EXPOSE 8080

CMD ["./app"]