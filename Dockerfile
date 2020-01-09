FROM golang:alpine

RUN apk add --no-cache git gcc musl-dev

RUN CGO_ENABLED=1 go get github.com/alash3al/scraply

ENTRYPOINT ["scraply"]

WORKDIR /root/