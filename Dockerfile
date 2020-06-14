FROM golang:1.14

WORKDIR /go/src/app
COPY . .

EXPOSE 80

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go", "run", "server.go"]