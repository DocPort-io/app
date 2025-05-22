import PocketBase from 'pocketbase';
import * as uuid from 'uuid';

export type TestUser = {
	id: string;
	email: string;
	password: string;
};

export const createUser = async (pocketBase: PocketBase): Promise<TestUser> => {
	const email = `test-${uuid.v4()}@example.com`;
	const password = 'password123';

	const { id } = await pocketBase.collection('users').create({
		name: 'Test User',
		email,
		password,
		passwordConfirm: password
	});

	return {
		id,
		email,
		password
	};
};

export const deleteUser = async (pocketBase: PocketBase, user: TestUser) => {
	await pocketBase.collection('users').delete(user.id);
};
