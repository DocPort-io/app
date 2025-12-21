# syntax=docker/dockerfile:1

FROM alpine:3.23 AS runtime

ARG TARGETPLATFORM

WORKDIR /app

RUN apk add --no-cache ca-certificates curl su-exec tzdata

# The GoReleaser docker pipe will provide the built binary named "app" in the build context.
COPY $TARGETPLATFORM/app /usr/bin/app
COPY ./scripts/docker /app/docker
COPY ./config.docker.toml /etc/docport/config.toml

RUN chmod +x /app/app \
    && find /app/docker -name "*.sh" -exec chmod +x {} \;

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 CMD curl -f http://localhost:8080/heartbeat || exit 1

ENTRYPOINT ["sh", "/app/docker/entrypoint.sh"]
CMD [ "/app/app" ]
