FROM golang:1.14 AS builder
WORKDIR /go/src/github.com/clnbs/wordChains
COPY . .
RUN go get -u ./...
RUN go mod vendor
RUN GO111MODULE=on go build -o astar.bin cmd/astar/main.go