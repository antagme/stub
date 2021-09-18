FROM golang:1.16-alpine AS builder

RUN apk update && apk add --no-cache \
    # SSL CA certificates are required to call HTTPS endpoints
    ca-certificates \

RUN mkdir -p /opt/stub
WORKDIR /opt/stub

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/stub
RUN ls -lah

FROM scratch

COPY --from=builder /opt/stub/bin/stub .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/stub"]
