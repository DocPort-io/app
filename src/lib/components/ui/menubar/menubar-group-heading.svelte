<script lang="ts">
	import type { ComponentProps } from 'svelte';

	import { cn } from '$lib/utils.js';
	import { Menubar as MenubarPrimitive } from 'bits-ui';

	let {
		ref = $bindable(null),
		inset,
		class: className,
		...restProps
	}: ComponentProps<typeof MenubarPrimitive.GroupHeading> & {
		inset?: boolean;
	} = $props();
</script>

<MenubarPrimitive.GroupHeading
	bind:ref
	data-slot="menubar-group-heading"
	data-inset={inset}
	class={cn('px-2 py-1.5 text-sm font-medium data-[inset]:pl-8', className)}
	{...restProps}
/>
