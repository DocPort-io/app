<script lang="ts">
	import { cn } from '$lib/utils.js';
	import { Menubar as MenubarPrimitive } from 'bits-ui';
	import { ChevronRight } from '@lucide/svelte';

	type $$Props = MenubarPrimitive.SubTriggerProps & {
		inset?: boolean;
	};
	type $$Events = MenubarPrimitive.SubTriggerEvents;

	let className: $$Props['class'] = undefined;
	export let inset: $$Props['inset'] = undefined;
	export { className as class };
</script>

<MenubarPrimitive.SubTrigger
	class={cn(
		'data-highlighted:bg-accent data-[state=open]:bg-accent data-highlighted:text-accent-foreground data-[state=open]:text-accent-foreground flex cursor-default items-center rounded-sm px-2 py-1.5 text-sm outline-hidden select-none data-disabled:pointer-events-none data-disabled:opacity-50',
		inset && 'pl-8',
		className
	)}
	on:click
	{...$$restProps}
	on:keydown
	on:focusin
	on:focusout
	on:pointerleave
	on:pointermove
>
	<slot />
	<ChevronRight class="ml-auto h-4 w-4" />
</MenubarPrimitive.SubTrigger>
