<script lang="ts">
	import { cn } from '$lib/utils.js';
	import { Drawer as DrawerPrimitive } from 'vaul-svelte';

	type $$Props = DrawerPrimitive.TitleProps;

	export let el: $$Props['el'] = undefined;
	let className: $$Props['class'] = undefined;
	export { className as class };
</script>

<DrawerPrimitive.Title
	bind:el
	class={cn('text-lg leading-none font-semibold tracking-tight', className)}
	{...$$restProps}
>
	<slot />
</DrawerPrimitive.Title>
