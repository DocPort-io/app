<script lang="ts">
	import type { Snippet } from 'svelte';

	import { goto } from '$app/navigation';
	import { page } from '$app/state';
	import { AppRoute } from '$lib/constants';
	import { setUserState } from '$lib/stores/user.svelte';

	type Props = {
		children: Snippet;
	};

	let { children }: Props = $props();

	const userState = setUserState();

	$effect(() => {
		if (userState.isValid) return;
		if (page.url.pathname === AppRoute.LOGIN()) return;

		const redirect = page.url.href.replace(page.url.origin, '');
		goto(`${AppRoute.LOGIN()}?redirect=${redirect}`);
	});

	$effect(() => {
		if (!userState.isValid) return;
		if (page.url.pathname !== AppRoute.LOGIN()) return;

		if (!page.url.searchParams.has('redirect')) {
			goto(AppRoute.DASHBOARD());
			return;
		}

		const redirect = page.url.searchParams.get('redirect');

		if (!redirect) {
			goto(AppRoute.DASHBOARD());
			return;
		}

		goto(redirect);
	});
</script>

{@render children()}
