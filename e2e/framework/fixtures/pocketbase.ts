import PocketBase from 'pocketbase';
import z from 'zod';

export const createPocketBase = async () => {
	const { TEST_POCKETBASE_URL, TEST_POCKETBASE_ADMIN_EMAIL, TEST_POCKETBASE_ADMIN_PASSWORD } =
		parseEnvironmentVariables();

	const pocketBase = new PocketBase(TEST_POCKETBASE_URL);

	await pocketBase
		.collection('_superusers')
		.authWithPassword(TEST_POCKETBASE_ADMIN_EMAIL, TEST_POCKETBASE_ADMIN_PASSWORD);

	return pocketBase;
};

const parseEnvironmentVariables = () => {
	return z
		.object({
			TEST_POCKETBASE_URL: z.string().url(),
			TEST_POCKETBASE_ADMIN_EMAIL: z.string(),
			TEST_POCKETBASE_ADMIN_PASSWORD: z.string()
		})
		.parse(process.env);
};
