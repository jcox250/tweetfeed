FROM golang:1.9

WORKDIR /go/src/demoServer

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["demoServer"]

