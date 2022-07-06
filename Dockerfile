FROM golang:1.18-alpine As builder

WORKDIR /scraply/

RUN apk update && apk add git upx

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o /usr/bin/scraply ./cmd/

RUN upx -9 /usr/bin/scraply

FROM scratch

WORKDIR /scraply/

COPY --from=builder /usr/bin/scraply /usr/bin/scraply
