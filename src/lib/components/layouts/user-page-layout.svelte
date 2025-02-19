<script lang="ts">
	import type { Snippet } from 'svelte';
	import NavigationBar from '../shared/navigation-bar.svelte';

	interface Props {
		hideNavbar?: boolean;
		children?: Snippet;
	}

	let { hideNavbar = false, children }: Props = $props();
</script>

<div class="flex min-h-screen w-full flex-col">
	{#if !hideNavbar}
		<NavigationBar />
	{/if}
	<main class="flex flex-1 flex-col gap-4 p-4 md:gap-8 md:p-8">
		{@render children?.()}
	  </main>
</div>
