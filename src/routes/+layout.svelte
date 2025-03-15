<script lang="ts">
	import '../app.css';

	import { dev } from '$app/environment';
	import { setAppState } from '$lib/stores/app.svelte';
	import { setProjects } from '$lib/stores/projects.svelte';
	import { setUserState } from '$lib/stores/user.svelte';
	import { ModeWatcher, setMode, mode } from 'mode-watcher';
	import { RenderScan } from 'svelte-render-scan';

	let { children } = $props();

	const appState = setAppState();
	setUserState();
	setProjects();

	$effect(() => {
		if (!appState.theme) appState.theme = $mode;
		if (appState.theme) setMode(appState.theme);
	});
</script>

{#if dev}
	<RenderScan />
{/if}
<ModeWatcher />
{@render children()}
