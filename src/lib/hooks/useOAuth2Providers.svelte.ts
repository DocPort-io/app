import type { AuthProviderInfo } from 'pocketbase';

import { getPocketBase, type TypedPocketBase } from '$lib/services/pocketbase';

export const useOAuth2Providers = (pocketbase: TypedPocketBase = getPocketBase()) => {
	const providers = $state<AuthProviderInfo[]>([]);
	let loading = $state(false);
	let error: string | null = $state(null);

	loading = true;
	pocketbase
		.collection('users')
		.listAuthMethods({ requestKey: null })
		.then((authMethods) => {
			providers.push(...authMethods.oauth2.providers);
			loading = false;
		})
		.catch((err) => {
			error = err instanceof Error ? err.message : err;
			loading = false;
		});

	return {
		get providers() {
			return providers;
		},
		get loading() {
			return loading;
		},
		get error() {
			return error;
		}
	};
};
