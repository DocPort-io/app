<script lang="ts">
	import { cn } from '$lib/utils.js';
	import { Menubar as MenubarPrimitive } from 'bits-ui';

	type $$Props = MenubarPrimitive.TriggerProps;
	type $$Events = MenubarPrimitive.TriggerEvents;

	let className: $$Props['class'] = undefined;
	export { className as class };
</script>

<MenubarPrimitive.Trigger
	class={cn(
		'data-highlighted:bg-accent data-[state=open]:bg-accent data-highlighted:text-accent-foreground data-[state=open]:text-accent-foreground flex cursor-default items-center rounded-sm px-3 py-1.5 text-sm font-medium outline-hidden select-none',
		className
	)}
	on:click
	on:keydown
	on:pointerenter
	{...$$restProps}
>
	<slot />
</MenubarPrimitive.Trigger>
