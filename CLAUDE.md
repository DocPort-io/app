# Claude Instructions for DocPort App Repository

## Overview

DocPort is a modern team and project management application built with SvelteKit and PocketBase. It features team collaboration, project management, file versioning, and user authentication with OAuth2 support.

## Commands

- **Build**: `pnpm run build`
- **Dev**: `pnpm run dev` (add `-- --open` to open in browser)
- **Lint**: `pnpm run lint`
- **Format**: `pnpm run format`
- **Typecheck**: `pnpm run check`
- **Test (all)**: `pnpm run test`
- **Unit tests**: `pnpm run test:unit`
- **Single unit test**: `pnpm run test:unit -- --test="test name"` or `pnpm run test:unit path/to/spec.ts`
- **E2E tests**: `pnpm run test:e2e`

## Architecture Overview

### Tech Stack
- **Frontend**: SvelteKit with Svelte 5, TailwindCSS v4
- **Backend**: PocketBase (Go-based BaaS)
- **Database**: SQLite (via PocketBase)
- **Deployment**: Docker containers with multi-stage builds
- **State Management**: TanStack Query + Custom Svelte 5 stores
- **Forms**: Custom form system built with Svelte 5 runes + Zod v4
- **UI Components**: shadcn-svelte with custom theming
- **Testing**: Vitest (unit) + Playwright (E2E)
- **Internationalization**: Paraglide.js with English/Dutch support

### Project Structure

```
src/
├── lib/
│   ├── components/
│   │   ├── ui/           # shadcn-svelte components
│   │   ├── layouts/      # Layout components  
│   │   ├── shared/       # Shared app components (navigation, etc.)
│   │   └── projects/     # Feature-specific components
│   ├── form/             # Custom form system
│   ├── queries/          # TanStack Query definitions
│   ├── services/         # Data layer (PocketBase integration)
│   ├── schemas/          # Zod validation schemas
│   ├── stores/           # Svelte stores for global state
│   ├── hooks/            # Custom hooks/composables
│   └── utils.ts          # Utility functions
├── routes/
│   ├── (user)/           # Authenticated routes
│   │   ├── dashboard/
│   │   ├── projects/
│   │   └── auth/
│   └── +layout.svelte    # Root layout
└── app.css               # TailwindCSS + custom variables
```

## Key Architectural Patterns

### 1. Custom Form System
- **Location**: `/src/lib/form/`
- **Features**: Type-safe forms with Svelte 5 runes, Zod v4 validation, field-level validation
- **Usage**: Replaces FormSnap/SuperForms with lighter, more intuitive API
- **Key Files**:
  - `form.svelte.ts` - Core form logic with reactive state management
  - `field.svelte` - Field wrapper component with validation
  - Components for labels, errors, descriptions

### 2. Data Layer Architecture
- **Services**: Abstract data operations (`/src/lib/services/`)
- **Queries**: TanStack Query integration (`/src/lib/queries/`)
- **Schemas**: Zod validation schemas (`/src/lib/schemas/`)
- **Pattern**: Service classes implement interfaces, queries handle caching/mutations
- **Example**: ProjectService → createUpdateProjectMutation → UI components

### 3. Component Organization
- **UI Components**: shadcn-svelte based, in `/src/lib/components/ui/`
- **Feature Components**: Organized by domain (`projects/`, `shared/`)
- **Layout Pattern**: Route groups with shared layouts
- **Naming**: PascalCase for components, kebab-case for files

### 4. State Management
- **Global State**: Custom Svelte 5 stores (user, team, dialogs)
- **Server State**: TanStack Query for API data
- **Form State**: Custom form system with reactive validation
- **Context Pattern**: Stores exposed via context (see `UserState`, `TeamState`)

### 5. Type Safety
- **Zod Integration**: Schema-first approach with runtime validation
- **TypeScript**: Strict mode, explicit types, generic constraints
- **PocketBase Types**: Custom typed client (`TypedPocketBase`)
- **Form Types**: Full type inference from Zod schemas

## Code Style Guidelines

### TypeScript & Code Quality
- **TypeScript**: Strict mode enabled; use explicit types
- **Imports**: Must be sorted (enforced by eslint-plugin-perfectionist)
- **Error Handling**: Use try/catch blocks with appropriate error typing
- **Naming**: Use camelCase for variables/functions, PascalCase for components

### Framework Patterns
- **Svelte 5**: Use new runes syntax (`$state`, `$derived`, `$effect`)
- **Components**: Follow single responsibility principle
- **Props**: Use `let { prop } = $props()` syntax
- **Reactivity**: Prefer `$derived` over `$effect` when possible

### Styling & UI
- **TailwindCSS**: v4 with custom design tokens in CSS variables
- **Components**: shadcn-svelte architecture with customizations
- **Theming**: CSS custom properties for consistent design system
- **Responsive**: Mobile-first approach

### Testing Strategy
- **Unit Tests**: Vitest with Testing Library for components
- **E2E Tests**: Playwright with custom test fixtures and utilities
- **Test Organization**: Tests co-located with features when possible
- **Mock Strategy**: Vi mocks for external dependencies

## Development Workflow

### Environment Setup
1. **Requirements**: Node.js 18+, pnpm, Docker
2. **Backend**: `docker compose -f compose.yml -f compose.dev.override.yml up backend`
3. **Frontend**: `pnpm run dev`
4. **Database**: PocketBase admin at http://localhost:8080/_/

### Database Schema
- **Collections**: users, teams, projects, versions, files
- **Relations**: Users belong to teams, projects belong to teams
- **Auth**: Built-in PocketBase authentication + OAuth2
- **Migrations**: JavaScript files in `/pb_migrations/`

### Form Development Pattern
1. Define Zod schema in `/src/lib/schemas/`
2. Create service methods in `/src/lib/services/`
3. Set up queries/mutations in `/src/lib/queries/`
4. Build form component using custom form system
5. Handle validation and submission with type safety

### Component Development
1. Use shadcn-svelte base components from `/src/lib/components/ui/`
2. Create feature-specific components in appropriate domain folders
3. Follow accessibility patterns (proper labeling, keyboard navigation)
4. Use TailwindCSS with design system tokens
5. Test with both unit and E2E tests

## Important Configuration

### Build & Deployment
- **Adapter**: `@sveltejs/adapter-node` for Docker deployment
- **SSR**: Disabled (`export const ssr = false`)
- **Docker**: Multi-stage builds with security-focused production images
- **Environment**: PocketBase URL configurable via `PUBLIC_POCKETBASE_URL`

### Internationalization
- **Tool**: Paraglide.js with compile-time optimization
- **Languages**: English (en) and Dutch (nl)
- **Usage**: Import `m` from paraglide messages, use in schemas and UI
- **Strategy**: localStorage + preferred language + base locale fallback

### Performance Optimizations
- **Query Caching**: TanStack Query with smart invalidation
- **Form Validation**: Conditional validation to prevent reactive loops
- **Bundle Size**: Tree-shaking, dynamic imports where appropriate
- **Images**: Optimized avatar handling through PocketBase file API
