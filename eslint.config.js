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

export default ts.config(
	includeIgnoreFile(gitignorePath),
	js.configs.recommended,
	...ts.configs.recommended,
	...svelte.configs['flat/recommended'],
	prettier,
	...svelte.configs['flat/prettier'],
	{
		languageOptions: {
			globals: {
				...globals.browser,
				...globals.node
			}
		}
	},
	{
		files: ['**/*.svelte'],

		languageOptions: {
			parserOptions: {
				parser: ts.parser
			}
		}
	},
	{
		files: ['**/*.{ts,svelte}'],
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
		...playwright.configs['flat/recommended'],
		files: ['e2e/**'],
		rules: {
			...playwright.configs['flat/recommended'].rules
		}
	}
);
