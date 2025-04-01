# Contributing to DocPort

Thank you for considering contributing to DocPort! This document provides guidelines and instructions for contributing to this project.

## Contact Information

For questions or concerns about this project, you can reach out to:

- Email: jonas@jonasclaes.be
- GitHub Issues: Create an issue in the repository

## Code of Conduct

Please read our [Code of Conduct](CODE_OF_CONDUCT.md) before participating in this project.

## How Can I Contribute?

### Reporting Bugs

Bugs are tracked as GitHub issues. Create an issue and provide the following information:

- Clear and descriptive title
- Steps to reproduce the behavior
- Expected behavior
- Actual behavior
- Screenshots (if applicable)
- Environment details (browser, OS, etc.)

### Suggesting Enhancements

Enhancement suggestions are also tracked as GitHub issues. When creating an enhancement suggestion, include:

- Clear and descriptive title
- Detailed description of the proposed functionality
- Explanation of why this enhancement would be useful
- Possible implementation approach (if you have ideas)

### Pull Requests

Follow these steps to submit a pull request:

1. Fork the repository
2. Create a new branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Run tests (`pnpm run test`)
5. Commit your changes (`git commit -m 'Add some amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

## Development Setup

1. Install dependencies:
   ```bash
   pnpm install
   ```

2. Start the development environment:
   ```bash
   docker compose -f compose.yml -f compose.dev.override.yml up
   pnpm run dev
   ```

3. Run tests to verify your changes:
   ```bash
   pnpm run test
   ```

## Styleguides

### Git Commit Messages

- Use the present tense ("Add feature" not "Added feature")
- Use the imperative mood ("Move cursor to..." not "Moves cursor to...")
- Limit the first line to 72 characters or less
- Reference issues and pull requests after the first line

### TypeScript Styleguide

- Use strict type checking
- Write clean, readable, and self-documenting code
- Follow the existing coding style (enforced by ESLint)
- Use explicit types rather than implicit ones

### Svelte Component Styleguide

- Follow the existing component structure
- Use TypeScript for all components
- Maintain responsive design principles
- Ensure accessibility compliance
- Follow shadcn-svelte patterns for UI components

## Additional Notes

### Issue and Pull Request Labels

We use labels to categorize issues and pull requests:

- `bug`: Something isn't working
- `enhancement`: New feature or request
- `documentation`: Improvements or additions to documentation
- `good first issue`: Good for newcomers
- `help wanted`: Extra attention is needed

Thank you for contributing to DocPort!

## Getting Help

If you need help with your contribution or have questions about the project, please feel free to:

1. Contact Jonas directly at jonas@jonasclaes.be
2. Open an issue with your question
3. Check existing documentation before reaching out