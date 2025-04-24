<script lang="ts">
	import { cn } from '$lib/utils.js';
	import { Menubar as MenubarPrimitive } from 'bits-ui';

	let {
		ref = $bindable(null),
		class: className,
		inset = undefined,
		...restProps
	}: MenubarPrimitive.ItemProps & {
		inset?: boolean;
	} = $props();
</script>

<MenubarPrimitive.Item
	bind:ref
	class={cn(
		'data-[highlighted]:bg-accent data-[highlighted]:text-accent-foreground relative flex cursor-default items-center rounded-sm px-2 py-1.5 text-sm outline-none select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50',
		inset && 'pl-8',
		className
	)}
	{...restProps}
/>
