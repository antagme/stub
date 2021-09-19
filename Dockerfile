FROM golang:1.16-alpine AS builder

RUN apk update && apk add --no-cache \
    # SSL CA certificates are required to call HTTPS endpoints
    ca-certificates \
    # Make build automation tool
    make \
    # Git required to operate with the git repository
    git \
    # Zoneinfo for timezones
    tzdata \
    # wget to download stuff
    wget

# A glibc compatibility layer package for Alpine Linux
# protoc is not statically compiled, and adding this package fixes the problem.
RUN wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub
RUN wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.29-r0/glibc-2.29-r0.apk
RUN apk add glibc-2.29-r0.apk

RUN mkdir -p /opt/stub
WORKDIR /opt/stub

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/stub
RUN ls -lah

FROM scratch

COPY --from=builder /opt/stub/bin/stub .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/stub"]
