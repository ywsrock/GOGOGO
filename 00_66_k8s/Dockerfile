FROM golang:1.22.0-alpine3.19 AS builder
WORKDIR /src/
COPY main.go go.* /src/

RUN CGO_ENABLED=0 go build -o /bin/demo

FROM scratch
COPY --from=builder /bin/demo /bin/demo
ENTRYPOINT ["/bin/demo"]