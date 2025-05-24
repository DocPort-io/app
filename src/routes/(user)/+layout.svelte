<script lang="ts">
	import type { Snippet } from 'svelte';

	import { createSyncStoragePersister } from '@tanstack/query-sync-storage-persister';
	import { QueryClient } from '@tanstack/svelte-query';
	import { SvelteQueryDevtools } from '@tanstack/svelte-query-devtools';
	import { PersistQueryClientProvider } from '@tanstack/svelte-query-persist-client';
	import { browser, dev } from '$app/environment';
	import TeamState from '$lib/components/shared/team-state.svelte';
	import UserState from '$lib/components/shared/user-state.svelte';
	import { Toaster } from '$lib/components/ui/sonner';
	import { ModeWatcher } from 'mode-watcher';
	import { RenderScan } from 'svelte-render-scan';

	type Props = {
		children: Snippet;
	};

	let { children }: Props = $props();

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
	<RenderScan initialEnabled={false} />
{/if}
<ModeWatcher />
<Toaster />
<PersistQueryClientProvider client={queryClient} persistOptions={{ persister }}>
	<UserState>
		<TeamState>
			{@render children()}
		</TeamState>
	</UserState>
	{#if dev}
		<SvelteQueryDevtools buttonPosition="bottom-left" />
	{/if}
</PersistQueryClientProvider>
