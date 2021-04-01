FROM golang:1.16.2 AS builder
WORKDIR /workdir
COPY ./ /workdir
RUN go build -ldflags "-linkmode external -extldflags -static" ./src/main.go

FROM scratch
COPY --from=builder /workdir/main /main
EXPOSE 8080
CMD ["/main"]
