<script lang="ts">
	import '../app.css';

	import { setAppState } from '$lib/stores/app.svelte';
	import { setProjects } from '$lib/stores/projects.svelte';
	import { setUserState } from '$lib/stores/user.svelte';
	import { ModeWatcher, setMode, mode } from 'mode-watcher';

	let { children } = $props();

	const appState = setAppState();
	setUserState();
	setProjects();

	$effect(() => {
		if (!appState.theme) appState.theme = $mode;
		if (appState.theme) setMode(appState.theme);
	});
</script>

<ModeWatcher />
{@render children()}
