FROM golang:1.9

WORKDIR /go/src/tweetfeed

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["tweetfeed"]

