FROM golang:alpine
RUN apk add --no-cache git

WORKDIR /go/src/github.com/otz1/otz
COPY . .

RUN go get ./...
RUN go install

CMD ["otz"]

EXPOSE 8001