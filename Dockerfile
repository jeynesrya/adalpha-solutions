
FROM golang:1.12.1

WORKDIR $GOPATH/src/github.com/jeynesrya/adalpha-solutions

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8080

CMD ["adalpha-solutions"]