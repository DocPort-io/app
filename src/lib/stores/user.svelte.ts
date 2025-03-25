import type { UserSchema } from '$lib/schemas/user.schema';
import type { AuthRecord } from 'pocketbase';

import { getPocketBase, type TypedPocketBase } from '$lib/services/pocketbase';
import { getContext, setContext } from 'svelte';

export class UserState {
	token = $state<string>('');
	name = $state<string>('');
	avatarUrl = $state<string | null>(null);

	constructor(protected readonly pocketbase: TypedPocketBase = getPocketBase()) {
		this.token = pocketbase.authStore.token;
		this.updateUserFromAuthRecord(pocketbase.authStore.record);

		pocketbase.authStore.onChange((token, record) => {
			this.token = token;
			this.updateUserFromAuthRecord(record);
		});
	}

	protected updateUserFromAuthRecord(record: AuthRecord) {
		if (!record) return;

		const authRecord = record as unknown as UserSchema;
		this.name = authRecord.name;
		this.avatarUrl = this.pocketbase.files.getURL(authRecord, authRecord.avatar) || null;
	}

	async signIn(email: string, password: string) {
		await this.pocketbase.collection('users').authWithPassword(email, password);
	}

	logout() {
		this.pocketbase.authStore.clear();
	}

	async getOAuth2Providers() {
		return (await this.pocketbase.collection('users').listAuthMethods()).oauth2.providers;
	}

	async signInWithExternalProvider(provider: string) {
		await this.pocketbase.collection('users').authWithOAuth2({
			provider
		});
	}
}

const USER_STATE_KEY = Symbol('USER_STATE');

export const setUserState = () => {
	return setContext(USER_STATE_KEY, new UserState());
};

export const getUserState = () => {
	return getContext<ReturnType<typeof setUserState>>(USER_STATE_KEY);
};
