FROM golang:1.24 as builder

LABEL org.opencontainers.image.source=https://github.com/nocturnalBadger/adjective-animal

WORKDIR /build
ADD main.go go.mod /build/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" -o /build/adjective-animal

FROM scratch

COPY --from=builder /build/adjective-animal /adjective-animal

ENTRYPOINT [ "/adjective-animal", "-server" ]
