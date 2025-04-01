# DocPort

DocPort is a modern team and project management application built with SvelteKit and PocketBase. It allows teams to organize projects, track progress, and collaborate efficiently.

## Features

- **Team Management**: Create and manage teams with different permissions
- **Project Management**: Create, update, and delete projects with status tracking
- **User Authentication**: Secure login system with OAuth 2.0 support
- **Responsive Design**: Beautiful UI that works on desktop and mobile
- **Internationalization**: Supports multiple languages including English and Dutch

## Getting Started

### Prerequisites

- Node.js v18+ and pnpm
- Docker and Docker Compose (for development with PocketBase)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/DocPort-io/app.git
   cd app
   ```

2. Install dependencies:
   ```bash
   pnpm install
   ```

3. Start the development environment:
   ```bash
   docker compose -f compose.yml -f compose.dev.override.yml up backend
   ```

4. Start the frontend development server:
   ```bash
   pnpm run dev
   ```

## Running the Tests

### Unit Tests

Run unit tests with Vitest:

```bash
pnpm run test:unit
```

To run a specific test:

```bash
pnpm run test:unit -- --test="test name"
# or
pnpm run test:unit path/to/spec.ts
```

### End-to-End Tests

Run E2E tests with Playwright:

```bash
pnpm run test:e2e
```

## Deployment

The application can be deployed using Docker:

1. Build the Docker containers (if not already built):
   ```bash
   docker compose build
   ```

2. Start the services in detached mode:
   ```bash
   docker compose up -d
   ```

This deploys both the frontend SvelteKit application and the PocketBase backend. The frontend will be available on port 3000 and the backend on port 8080 by default.

## Built With

- [SvelteKit](https://kit.svelte.dev/) - Web framework
- [Svelte 5](https://svelte.dev/) - Component framework
- [PocketBase](https://pocketbase.io/) - Backend database and auth
- [TailwindCSS](https://tailwindcss.com/) - CSS framework
- [Shadcn-Svelte](https://www.shadcn-svelte.com/) - UI components
- [Docker](https://www.docker.com/) - Containerization

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/DocPort-io/app/tags) and the [GitHub Releases page](https://github.com/DocPort-io/app/releases).

## Authors

- **Jonas Claes** - *Lead Engineer* - [GitHub](https://github.com/jonasclaes) - [Email](mailto:jonas@jonasclaes.be)

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- SvelteKit team for the excellent framework
- PocketBase for the lightweight backend solution
- Shadcn for the UI component inspiration
