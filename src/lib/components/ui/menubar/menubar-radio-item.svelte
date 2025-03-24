<script lang="ts">
	import { cn } from '$lib/utils.js';
	import { Menubar as MenubarPrimitive } from 'bits-ui';
	import { Circle } from '@lucide/svelte';

	type $$Props = MenubarPrimitive.RadioItemProps;
	type $$Events = MenubarPrimitive.RadioItemEvents;

	let className: $$Props['class'] = undefined;
	export let value: $$Props['value'];
	export { className as class };
</script>

<MenubarPrimitive.RadioItem
	{value}
	class={cn(
		'data-highlighted:bg-accent data-highlighted:text-accent-foreground relative flex cursor-default items-center rounded-sm py-1.5 pr-2 pl-8 text-sm outline-hidden select-none data-disabled:pointer-events-none data-disabled:opacity-50',
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
	<span class="absolute left-2 flex h-3.5 w-3.5 items-center justify-center">
		<MenubarPrimitive.RadioIndicator>
			<Circle class="h-2 w-2 fill-current" />
		</MenubarPrimitive.RadioIndicator>
	</span>
	<slot />
</MenubarPrimitive.RadioItem>
