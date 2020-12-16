FROM golang:1.14 AS tester
WORKDIR /go/src/github.com/clnbs/wordChains
RUN go get golang.org/x/tools/cmd/cover
COPY . .
RUN go get -u ./...
RUN go mod vendor
RUN go test -v -coverprofile cover.out ./...
RUN go tool cover -func=cover.out
RUN go tool cover -html=cover.out -o cover.html