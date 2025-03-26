<script lang="ts">
	import type { Snippet } from 'svelte';

	import NavigationBar from '../shared/navigation-bar.svelte';

	interface Props {
		title?: string;
		hideNavbar?: boolean;
		children?: Snippet;
	}

	let { title, hideNavbar = false, children }: Props = $props();
</script>

<svelte:head>
	<title>{`DocPort - ${title || 'App'}`}</title>
</svelte:head>

<div class="flex min-h-screen w-full flex-col">
	{#if !hideNavbar}
		<NavigationBar />
	{/if}
	<main class="flex flex-1 flex-col gap-2 p-4 md:gap-4 md:p-8 md:pt-4">
		{@render children?.()}
	</main>
</div>
