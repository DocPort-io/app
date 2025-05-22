import { defineConfig, devices } from '@playwright/test';
import 'dotenv/config';

export default defineConfig({
	webServer: {
		command: 'pnpm run build && pnpm run preview',
		port: 5173,
		reuseExistingServer: !process.env.CI
	},

	testDir: 'e2e',
	fullyParallel: true,

	projects: [
		{
			name: 'chromium',
			...devices['Desktop Chrome']
		}
	]
});
