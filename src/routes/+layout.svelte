<script lang="ts">
	import '../app.css';

	import { dev } from '$app/environment';
	import { goto } from '$app/navigation';
	import { page } from '$app/state';
	import { setAppState } from '$lib/stores/app.svelte';
	import { setProjects } from '$lib/stores/projects.svelte';
	import { setUserState } from '$lib/stores/user.svelte';
	import { ModeWatcher, setMode, mode } from 'mode-watcher';
	import { RenderScan } from 'svelte-render-scan';

	let { children } = $props();

	const appState = setAppState();
	const userState = setUserState();
	setProjects();

	$effect(() => {
		if (!appState.theme) appState.theme = $mode;
		if (appState.theme) setMode(appState.theme);
	});

	$effect(() => {
		if (userState.token !== '') return;
		if (page.url.pathname === '/auth/login') return;

		const redirect = page.url.href.replace(page.url.origin, '');
		goto(`/auth/login?redirect=${redirect}`);
	});

	$effect(() => {
		if (userState.token === '') return;
		if (page.url.searchParams.has('redirect')) {
			const redirect = page.url.searchParams.get('redirect');
			if (!redirect) return;
			goto(redirect);
		}
	});
</script>

{#if dev}
	<RenderScan />
{/if}
<ModeWatcher />
{@render children()}
