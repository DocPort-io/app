<script lang="ts">
	import '../app.css';

	import { dev } from '$app/environment';
	import { goto } from '$app/navigation';
	import { page } from '$app/state';
	import { Toaster } from '$lib/components/ui/sonner';
	import { AppRoute } from '$lib/constants';
	import { setAppState } from '$lib/stores/app.svelte';
	import { setProjects } from '$lib/stores/projects.svelte';
	import { setTeamState } from '$lib/stores/team.svelte';
	import { setUserState } from '$lib/stores/user.svelte';
	import { ModeWatcher, setMode, mode } from 'mode-watcher';
	import { RenderScan } from 'svelte-render-scan';

	let { children } = $props();

	const appState = setAppState();
	const userState = setUserState();
	const teamState = setTeamState();
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
		if (page.url.pathname !== '/auth/login') return;

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
		if (userState.token === '') return;
		teamState.load();
	});
</script>

{#if dev}
	<RenderScan />
{/if}
<ModeWatcher />
<Toaster />
{@render children()}
