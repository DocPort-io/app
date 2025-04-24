<script lang="ts">
	import { cn } from '$lib/utils.js';
	import { Menubar as MenubarPrimitive } from 'bits-ui';

	let {
		ref = $bindable(null),
		class: className,
		...restProps
	}: MenubarPrimitive.SubContentProps = $props();
</script>

<MenubarPrimitive.SubContent
	bind:ref
	class={cn(
		'bg-popover text-popover-foreground z-50 min-w-max rounded-md border p-1 focus:outline-none',
		className
	)}
	{...restProps}
/>
