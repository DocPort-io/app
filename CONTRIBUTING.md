# Contributing to DocPort.io

Thank you for your interest in contributing! This guide helps you set up a local environment, understand the workflow, and submit high‑quality changes.

## Code of Conduct

Be respectful and constructive. We aim for an inclusive, collaborative environment.

## Project Overview

DocPort.io is an API‑first backend in Go for managing technical documentation tied to physical equipment via QR codes. Learn more at https://docport.io and see the top‑level README for architecture and usage.

## Getting Started

### Prerequisites

- Go 1.21+

### Setup

```
git clone https://github.com/docport-io/app.git
cd app
cp config.example.toml config.toml
go mod download
go run ./cmd/app
```

### Configuration

- Adjust config.toml as needed (server bind/port, database URL, storage provider)
- Override via env vars with DOCPORT_ prefix (dots become underscores), e.g. DOCPORT_SERVER_PORT=8080

## Development Workflow

### Branching

- main is protected; create feature branches: feature/short-description, fix/short-description, chore/short-description

### Commits

- Prefer Conventional Commits style:
  - feat: add version tagging to file uploads
  - fix: correct pagination defaults for files list
  - chore: bump dependencies
  - docs: update README with API paths

### Coding Standards

- Language: Go
- Style: go fmt/go vet; keep imports organized
- Errors: wrap with context when helpful; return typed errors from service layer when applicable
- HTTP: handlers live in pkg/controller; keep business logic in pkg/service
- DTOs: in pkg/dto; avoid leaking DB models to controllers

### Database & SQL

- SQL is defined in query.sql (and migrations under migrations/)
- Generated code lives in pkg/database via sqlc
- If you change SQL queries or schema:
  1. Update migrations in migrations/*.up.sql and *.down.sql (use next number)
  2. Run sqlc generate to refresh pkg/database
  3. Start the app to apply migrations automatically

### Storage

- Default provider is filesystem. New providers should implement pkg/storage.FileStorage and be wired in pkg/app/storage.go

## Running and Testing

```
# run
go run ./cmd/app

# all tests
go test ./...
```

## API Documentation

- Swagger is served at /swagger (host configured via server.host). Update specs in pkg/docs if you change routes.

## Performance & Reliability

- Avoid blocking calls in handlers; defer heavy I/O to services
- Use pagination for list endpoints; validate inputs

## Submitting a Pull Request

### Checklist

- [ ] Branch created from latest main and rebased
- [ ] Code formatted (go fmt), vet clean (go vet)
- [ ] Tests added/updated and passing (go test ./...)
- [ ] Migrations included (if schema changes) and sqlc generated
- [ ] README/CONTRIBUTING updated if behavior or setup changed
- [ ] Swagger docs updated if API changed

### PR Review Guidelines

- Keep PRs scoped and < 400 lines when possible
- Provide context: what/why, screenshots for user‑visible changes
- Respond to review comments within 3 business days

## Issue Reporting

- Use clear titles; include steps to reproduce, expected vs actual, logs if relevant
- Label appropriately (bug, enhancement, docs)

## Security Policy

- Do not open public issues for vulnerabilities
- Report privately to the maintainers; request a security contact if not listed in repository settings

## License

By contributing, you agree that your contributions will be licensed under the repository’s license.
