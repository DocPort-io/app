# syntax=docker/dockerfile:1

FROM alpine:3.20 AS runtime

ARG TARGETPLATFORM

RUN apk add --no-cache ca-certificates tzdata \
    && adduser -D -s /sbin/nologin appuser \
    && mkdir /app \
    && chown appuser /app -R \
    && mkdir /app/storage

WORKDIR /app

COPY config.example.toml /etc/docport/config.toml

# The GoReleaser docker pipe will provide the built binary named "app" in the build context.
COPY $TARGETPLATFORM/app /app/app

EXPOSE 8080
USER appuser
ENTRYPOINT ["/app/app"]
