<script lang="ts">
	import { cn } from '$lib/utils.js';
	import { ContextMenu as ContextMenuPrimitive } from 'bits-ui';
	import { Check } from '@lucide/svelte';

	type $$Props = ContextMenuPrimitive.CheckboxItemProps;
	type $$Events = ContextMenuPrimitive.CheckboxItemEvents;

	let className: $$Props['class'] = undefined;
	export let checked: $$Props['checked'] = undefined;
	export { className as class };
</script>

<ContextMenuPrimitive.CheckboxItem
	bind:checked
	class={cn(
		'data-highlighted:bg-accent data-highlighted:text-accent-foreground relative flex cursor-default items-center rounded-sm py-1.5 pr-2 pl-8 text-sm outline-hidden select-none data-disabled:pointer-events-none data-disabled:opacity-50',
		className
	)}
	{...$$restProps}
	on:click
	on:keydown
	on:focusin
	on:focusout
	on:pointerdown
	on:pointerleave
	on:pointermove
>
	<span class="absolute left-2 flex h-3.5 w-3.5 items-center justify-center">
		<ContextMenuPrimitive.CheckboxIndicator>
			<Check class="h-4 w-4" />
		</ContextMenuPrimitive.CheckboxIndicator>
	</span>
	<slot />
</ContextMenuPrimitive.CheckboxItem>
