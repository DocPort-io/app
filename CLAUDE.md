# Claude Instructions for App Repository

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

## Code Style Guidelines

- **TypeScript**: Strict mode enabled; use explicit types
- **Imports**: Must be sorted (enforced by eslint-plugin-perfectionist)
- **Framework**: SvelteKit with Svelte 5
- **Testing**: Vitest for unit tests, Playwright for E2E tests
- **Components**: Follow shadcn-svelte architecture in UI components
- **Formatting**: Prettier with tailwind plugin
- **Naming**: Use camelCase for variables/functions, PascalCase for components
- **Error Handling**: Use try/catch blocks with appropriate error typing
- **State Management**: Prefer Svelte stores for shared state
