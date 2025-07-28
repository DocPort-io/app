import { includeIgnoreFile } from '@eslint/compat';
import js from '@eslint/js';
import prettier from 'eslint-config-prettier';
import perfectionist from 'eslint-plugin-perfectionist';
import playwright from 'eslint-plugin-playwright';
import svelte from 'eslint-plugin-svelte';
import globals from 'globals';
import { fileURLToPath } from 'node:url';
import ts from 'typescript-eslint';
const gitignorePath = fileURLToPath(new URL('./.gitignore', import.meta.url));
import svelteConfig from './svelte.config.js';

export default ts.config(
	includeIgnoreFile(gitignorePath),
	js.configs.recommended,
	...ts.configs.recommended,
	...svelte.configs.recommended,
	prettier,
	...svelte.configs.prettier,
	{
		languageOptions: {
			globals: {
				...globals.browser,
				...globals.node
			}
		}
	},
	{
		files: ['**/*.svelte', '**/*.svelte.ts', '**/*.svelte.js'],
		languageOptions: {
			parserOptions: {
				projectService: true,
				extraFileExtensions: ['.svelte'],
				parser: ts.parser,
				svelteConfig
			}
		}
	},
	{
		files: ['**/*.ts', '**/*.svelte'],
		plugins: {
			perfectionist
		},
		rules: {
			'perfectionist/sort-imports': 'error'
		}
	},
	{
		files: ['pb_hooks/**/*.js', 'pb_migrations/**/*.js'],
		rules: {
			'@typescript-eslint/triple-slash-reference': 'off',
			'no-undef': 'off'
		}
	},
	{
		files: ['e2e/**'],
		...playwright.configs['flat/recommended'],
		rules: {
			...playwright.configs['flat/recommended'].rules
		}
	}
);
