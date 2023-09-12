FROM golang:1.14.6-alpine3.12 as builder
COPY go.mod go.sum /go/src/github.com/murilo-toddy/mess/
WORKDIR /go/src/github.com/murilo-toddy/mess/
RUN go mod download
COPY . /go/src/github.com/murilo-toddy/mess/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/mess github.com/murilo-toddy/mess

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/murilo-toddy/mess/build/mess /usr/bin/mess
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/mess"]
