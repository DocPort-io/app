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

<div class="bg-muted/40 flex min-h-screen w-full flex-col">
	{#if !hideNavbar}
		<NavigationBar />
	{/if}
	<div class="flex flex-col sm:gap-4 sm:py-4">
		<main class="grid flex-1 items-start gap-4 p-4 sm:px-6 sm:py-0 md:gap-8">
			{@render children?.()}
		</main>
	</div>
</div>
