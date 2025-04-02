import { defineConfig, devices } from '@playwright/test';
import 'dotenv/config';

export default defineConfig({
	webServer: {
		command: 'pnpm run build && pnpm run preview',
		port: 4173,
		reuseExistingServer: !process.env.CI
	},

	testDir: 'e2e',

	projects: [
		{
			name: 'chromium',
			...devices['Desktop Chrome']
		}
	]
});
