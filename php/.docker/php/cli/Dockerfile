FROM ghcr.io/kuaukutsu/php:8.4-cli AS app_setup

# Arguments
ARG UID=10001
ARG WORKDIR="/src"

# Configure
COPY conf/php.ini /usr/local/etc/php/php.ini
# Dependencies
COPY --from=composer:latest --link /usr/bin/composer /usr/bin/composer
# user
RUN <<EOF
    adduser -u ${UID} -G www-data -s /bin/sh -D developer www-data
EOF
# workdir
RUN <<EOF
    mkdir -p ${WORKDIR}
    chown -R ${UID}:www-data ${WORKDIR}
EOF
# composer
RUN <<EOF
    mkdir /var/.composer
    chown developer:www-data /var/.composer
EOF
ENV COMPOSER_CACHE_DIR=/var/.composer

RUN install-php-extensions bcmath zlib

FROM app_setup AS app_protoc

ARG PROTOC_PLUGIN_VERSION

RUN apk add --no-cache protoc protobuf-dev

RUN curl -L "https://github.com/thesis-php/protoc-plugin/releases/download/${PROTOC_PLUGIN_VERSION}/protoc-gen-php" \
    -o /usr/local/bin/protoc-gen-php \
    && chmod +x /usr/local/bin/protoc-gen-php

FROM app_setup AS app_devel

RUN install-php-extensions xdebug

USER $UID
WORKDIR $WORKDIR

FROM app_setup AS app_cli

USER $UID
WORKDIR $WORKDIR
