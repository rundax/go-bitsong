FROM golang:1.13.15-alpine AS build-env

# Set up dependencies
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3

RUN apk add --no-cache $PACKAGES

# Set working directory for the build
WORKDIR /app

# Add source files
COPY . .

# Install minimum necessary dependencies, build Cosmos SDK, remove packages
RUN make go-mod-cache && \
    make build

# Final image
FROM alpine:latest

# Install ca-certificates
RUN apk add --update ca-certificates

## Copy over binaries from the build-env
COPY --from=build-env /app/build/bitsongd /usr/local/bin/bitsongd
COPY --from=build-env /app/build/bitsongcli /usr/local/bin/bitsongcli
##
#RUN apk add --no-cache shadow
RUN apk add bash
#
####
#### Envs
####
ENV MY_USER="bitsong" \
    MY_GROUP="bitsong" \
    MY_UID="1000" \
    MY_GID="1000"

####
#### User/Group
####
RUN set -eux \
    && addgroup -g ${MY_GID} ${MY_GROUP} \
    && adduser -u ${MY_UID} -G ${MY_GROUP} -s /bin/sh -D ${MY_USER}

# Must set this value for the bash shell to source
# the '/etc/bashrc' file.
# See: https://stackoverflow.com/q/29021704
ENV BASH_ENV /etc/bashrc
COPY ./.docker/bashrc /etc/bashrc


COPY ./.docker/docker-entrypoint.sh /docker-entrypoint.sh

RUN chmod +x /docker-entrypoint.sh

USER bitsong
WORKDIR /home/bitsong

ENTRYPOINT ["/docker-entrypoint.sh"]
