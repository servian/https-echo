#
# Builder
#
FROM golang:alpine AS builder
LABEL maintainer="Tristan Morgan <tristan@vibrato.com.au>"

WORKDIR /go/src/github.com/servian/https-echo/

COPY . .

ARG LD_FLAGS="-s -w"
ENV LD_FLAGS="${LD_FLAGS}"

RUN \
  CGO_ENABLED="0" \
  GOOS="linux" \
  GOARCH="amd64" \
  go build -a -o "/https-echo" -ldflags "${LD_FLAGS}"

#
# Final
#
FROM scratch
LABEL maintainer="Tristan Morgan <tristan@vibrato.com.au>"
LABEL Description="HTTPS_ECHO, echo url and redirect http to https"
EXPOSE 80
WORKDIR /
COPY --from=builder /https-echo /https-echo
ENTRYPOINT ["/https-echo"]
