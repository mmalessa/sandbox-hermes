FROM golang:1.24-alpine AS godev

RUN apk update && apk upgrade
RUN apk add --no-cache bash git openssh autoconf automake libtool gettext gettext-dev make g++ texinfo curl grpc-plugins

RUN   mkdir /go/pkg
RUN   chmod a+rwx /go/pkg

ARG DEVELOPER_UID=1000
RUN adduser -s /bin/sh -u ${DEVELOPER_UID} -D developer
USER developer

WORKDIR /go/src/hermes

FROM php:8.4-cli-alpine AS phpdev
RUN apk --no-cache add \
    unzip librdkafka tzdata protobuf grpc-plugins \
    $PHPIZE_DEPS linux-headers htop procps bash vim \
    && rm -rf /var/cache/apk*
RUN apk add grpcurl --repository=http://dl-cdn.alpinelinux.org/alpine/edge/testing

COPY --from=mlocati/php-extension-installer /usr/bin/install-php-extensions /usr/local/bin/
RUN set -eux;  \
    install-php-extensions zip pcntl intl bcmath pdo pdo_pgsql rdkafka ds sockets amqp grpc \
    && rm -rf /tmp/*
#ENV COMPOSER_ALLOW_SUPERUSER=1
ENV COMPOSER_HOME=/.composer
ENV TZ=UTC
COPY --from=composer/composer:2-bin /composer /usr/bin/composer

ARG DEVELOPER_UID=1000
RUN adduser -s /bin/sh -u ${DEVELOPER_UID} -D developer
USER developer
WORKDIR /app
