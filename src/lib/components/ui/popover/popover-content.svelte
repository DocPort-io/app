<script lang="ts">
	import { cn, flyAndScale } from '$lib/utils.js';
	import { Popover as PopoverPrimitive } from 'bits-ui';

	type $$Props = PopoverPrimitive.ContentProps;
	let className: $$Props['class'] = undefined;
	export let transition: $$Props['transition'] = flyAndScale;
	export let transitionConfig: $$Props['transitionConfig'] = undefined;
	export { className as class };
</script>

<PopoverPrimitive.Content
	{transition}
	{transitionConfig}
	class={cn(
		'bg-popover text-popover-foreground z-50 w-72 rounded-md border p-4 shadow-md outline-hidden',
		className
	)}
	{...$$restProps}
>
	<slot />
</PopoverPrimitive.Content>
