<script lang="ts">
	import '../app.css';
	import { ModeWatcher, setMode, mode } from 'mode-watcher';
	import { setAppState } from '$lib/states/app.svelte';
	import { setUserState } from '$lib/states/user.svelte';
	import { setProjectsState } from '$lib/states/projects.svelte';
	let { children } = $props();

	const appState = setAppState();
	setUserState();
	setProjectsState();

	$effect(() => {
		if (!appState.theme) appState.theme = $mode;
		if (appState.theme) setMode(appState.theme);
	});
</script>

<ModeWatcher />
{@render children()}
