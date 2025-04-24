<script lang="ts">
	import { cn } from '$lib/utils.js';
	import { Menubar as MenubarPrimitive } from 'bits-ui';

	let {
		ref = $bindable(null),
		class: className,
		...restProps
	}: MenubarPrimitive.RootProps = $props();
</script>

<MenubarPrimitive.Root
	bind:ref
	class={cn('bg-background flex h-10 items-center space-x-1 rounded-md border p-1', className)}
	{...restProps}
/>
