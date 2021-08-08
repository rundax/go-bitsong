# Simple usage with a mounted data directory:
# > docker build -t bitsong .
# > docker run -it -p 46657:46657 -p 46656:46656 -v ~/.bitsongd:/root/.bitsongd -v ~/.bitsongcli:/root/.bitsongcli bitsong bitsongd init
# > docker run -it -p 46657:46657 -p 46656:46656 -v ~/.bitsongd:/root/.bitsongd -v ~/.bitsongcli:/root/.bitsongcli bitsong bitsongd start
FROM golang:alpine AS build-env

# Set up dependencies
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python

# Set working directory for the build
WORKDIR /go/src/github.com/bitsongofficial/go-bitsong

# Add source files
COPY . .

# Install minimum necessary dependencies, build Cosmos SDK, remove packages
RUN apk add --no-cache $PACKAGES
RUN make go-mod-cache && \
    make install

# Final image
FROM alpine:edge

# Install ca-certificates
RUN apk add --update ca-certificates
WORKDIR /root

# Copy over binaries from the build-env
COPY --from=build-env /go/bin/bitsongd /usr/bin/bitsongd
COPY --from=build-env /go/bin/bitsongcli /usr/bin/bitsongcli

# Run bitsongd by default, omit entrypoint to ease using container with bitsongcli
CMD ["bitsongd"]