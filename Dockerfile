# syntax=docker/dockerfile:1

FROM cgr.dev/chainguard/static:latest AS runtime

ARG TARGETPLATFORM

ENV HOME /home/nonroot
WORKDIR /home/nonroot

COPY config.example.toml /etc/docport/config.toml

# The GoReleaser docker pipe will provide the built binary named "app" in the build context.
COPY $TARGETPLATFORM/app /app

EXPOSE 8080

ENTRYPOINT ["/app"]
