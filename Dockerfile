# syntax=docker/dockerfile:1

FROM alpine:3.20 AS runtime
RUN apk add --no-cache ca-certificates tzdata \
    && adduser -D -H -s /sbin/nologin appuser \
RUN mkdir /home/appuser
WORKDIR /app

COPY config.example.toml /etc/docport/config.toml

# The GoReleaser docker pipe will provide the built binary named "app" in the build context.
COPY app /app/app

EXPOSE 8080
USER appuser
ENTRYPOINT ["/app/app"]
