# DocPort.io — Modern Document Management for the Field

## Overview

DocPort.io is a modern document management platform that connects physical equipment to digital documentation using QR codes. It delivers version control, team collaboration, and field‑ready access to technical docs — even for large files — through a simple service.

## Key Features

- QR‑code anchored access to documentation for assets in the field
- Versioned documents with attach/detach to versions and projects
- File storage abstraction with local filesystem provider
- REST API with OpenAPI/Swagger docs available at /swagger
- Built‑in pagination middleware for list endpoints
- Zero‑downtime schema migrations on startup
- Configuration via file and environment variables

## Project Status

Early stage. API and storage are functional; expect breaking changes before v1.0.

## Links

- Website: https://docport.io
- Swagger UI: http(s)://{host}:{port}/swagger

## Architecture at a Glance

- Language: Go
- HTTP: chi router with middleware (request ID, logging, recover)
- Config: Viper (TOML + env)
- DB: SQLite (modernc.org/sqlite) by default
- Migrations: golang‑migrate with embedded SQL in the binary
- SQL: generated via sqlc (see sqlc.yaml and pkg/database)
- Storage: pluggable providers (filesystem included)

## Repository Layout

- cmd/app: application entrypoint
- pkg/app: server/bootstrap (routes, server, database, storage)
- pkg/controller: HTTP handlers
- pkg/service: business logic
- pkg/database: sqlc generated code and helpers
- migrations: SQL migrations (embedded into the binary)
- pkg/docs: Swagger/OpenAPI spec and generated bindings
- pkg/storage: storage provider interfaces/impls
- pkg/dto: request/response DTOs

## Quick Start

### Prerequisites

- Go 1.21+ (module mode)

1) Clone and enter

```
git clone https://github.com/docport-io/app.git
cd app
```

2) Configure

Copy the example config and adjust as needed:

```
cp config.example.toml config.toml
```

Important settings in config.toml:

- server.bind: interface to bind (e.g., 0.0.0.0)
- server.port: port (e.g., 8080)
- server.host: advertised host for Swagger (e.g., localhost:8080)
- database.driver: sqlite
- database.url: file:./test.db?cache=shared
- storage.provider: filesystem

### Environment variables

All config keys can be overridden via environment variables with prefix DOCPORT_ and dots replaced by underscores. Examples:

```
DOCPORT_SERVER.BIND=0.0.0.0     → DOCPORT_SERVER_BIND
DOCPORT_SERVER.PORT=8080        → DOCPORT_SERVER_PORT
DOCPORT_STORAGE.PROVIDER=...    → DOCPORT_STORAGE_PROVIDER
```

### Run

```
go run ./cmd/app
```

On first start, migrations run automatically and create/update the schema.

### Build a binary

```
go build -o docport ./cmd/app
./docport
```

### Storage Providers

Current provider: filesystem (stores uploads under ./storage)

Switch the provider by setting storage.provider in config.toml. Custom providers can implement pkg/storage.FileStorage and be wired in pkg/app/storage.go.

### Database & Migrations

- SQLite by default (embedded driver); connection string from config.toml
- Migrations are embedded (see embed.go and migrations/*) and applied on startup via golang‑migrate

### Development

#### Useful commands

```
# run
go run ./cmd/app

# test
go test ./...

# (if you modify SQL) regenerate with sqlc
sqlc generate
```

#### Testing

There are controller tests under pkg/controller. Run all tests:

```
go test ./...
```

## Contributing

See CONTRIBUTE.md for guidelines, environment setup, and the PR checklist.

## Security

Please report vulnerabilities responsibly. See CONTRIBUTE.md for the security policy and contact.

## License

Unless noted otherwise, this project is licensed under the Apache 2.0 License. See LICENSE or consult repository metadata.
