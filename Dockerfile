# syntax=docker/dockerfile:1

FROM alpine:3.20 AS runtime

ARG TARGETPLATFORM

RUN apk add --no-cache ca-certificates tzdata \
    && adduser -D -s /sbin/nologin appuser \
    && mkdir /app \
    && mkdir /app/storage \
    && chown appuser /app -R \
    && chmod u+rwx /app -R

WORKDIR /app

COPY config.docker.toml /etc/docport/config.toml

# The GoReleaser docker pipe will provide the built binary named "app" in the build context.
COPY $TARGETPLATFORM/app /usr/bin/app

EXPOSE 8080
USER appuser

ENTRYPOINT ["/usr/bin/app"]
