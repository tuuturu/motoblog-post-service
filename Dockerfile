FROM golang:1.17 AS build

ENV CGO_ENABLED=0

WORKDIR /go/src

COPY pkg ./pkg
COPY main.go .
COPY go.* ./

RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o service .

FROM scratch AS runtime

ENV GIN_MODE=release
EXPOSE 3000/tcp
ENTRYPOINT ["./service"]

COPY --from=build /go/src/service ./

