<script lang="ts">
	import '../../app.css';

	import { createSyncStoragePersister } from '@tanstack/query-sync-storage-persister';
	import { QueryClient } from '@tanstack/svelte-query';
	import { SvelteQueryDevtools } from '@tanstack/svelte-query-devtools';
	import { PersistQueryClientProvider } from '@tanstack/svelte-query-persist-client';
	import { browser, dev } from '$app/environment';
	import { goto } from '$app/navigation';
	import { page } from '$app/state';
	import { Toaster } from '$lib/components/ui/sonner';
	import { AppRoute } from '$lib/constants';
	import { setTeamState } from '$lib/stores/team.svelte';
	import { setUserState } from '$lib/stores/user.svelte';
	import { ModeWatcher } from 'mode-watcher';
	import { RenderScan } from 'svelte-render-scan';

	let { children } = $props();

	const userState = setUserState();
	setTeamState();

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

	$effect(() => {
		if (userState.isValid) return;
		goto(AppRoute.LOGIN());
	});

	const queryClient = new QueryClient({
		defaultOptions: {
			queries: {
				enabled: browser
			}
		}
	});

	const persister = createSyncStoragePersister({
		storage: browser ? window.localStorage : null
	});
</script>

{#if dev}
	<RenderScan />
{/if}
<ModeWatcher />
<Toaster />
<PersistQueryClientProvider client={queryClient} persistOptions={{ persister }}>
	{@render children()}
	{#if dev}
		<SvelteQueryDevtools initialIsOpen={true} />
	{/if}
</PersistQueryClientProvider>
