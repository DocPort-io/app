<script lang="ts">
	import { cn } from '$lib/utils.js';
	import { Menubar as MenubarPrimitive } from 'bits-ui';

	type $$Props = MenubarPrimitive.ItemProps & {
		inset?: boolean;
	};
	type $$Events = MenubarPrimitive.ItemEvents;

	let className: $$Props['class'] = undefined;
	export let inset: $$Props['inset'] = undefined;
	export { className as class };
</script>

<MenubarPrimitive.Item
	class={cn(
		'data-highlighted:bg-accent data-highlighted:text-accent-foreground relative flex cursor-default items-center rounded-sm px-2 py-1.5 text-sm outline-hidden select-none data-disabled:pointer-events-none data-disabled:opacity-50',
		inset && 'pl-8',
		className
	)}
	{...$$restProps}
	on:click
	on:keydown
	on:focusin
	on:focusout
	on:pointerleave
	on:pointermove
	on:pointerdown
>
	<slot />
</MenubarPrimitive.Item>
