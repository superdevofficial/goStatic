# stage 0
FROM golang:latest as builder
WORKDIR /go
COPY . .

RUN go get github.com/gorilla/mux
RUN GOARCH=amd64 GOOS=linux go build -o gostatic  -ldflags "-linkmode external -extldflags -static -w"

# stage 1
FROM centurylink/ca-certs
WORKDIR /
COPY --from=builder /go .
ENTRYPOINT ["/gostatic","-dir","/app"]
