<script lang="ts">
	import { cn } from '$lib/utils.js';
	import { Tabs as TabsPrimitive } from 'bits-ui';

	type $$Props = TabsPrimitive.TriggerProps;
	type $$Events = TabsPrimitive.TriggerEvents;

	let className: $$Props['class'] = undefined;
	export let value: $$Props['value'];
	export { className as class };
</script>

<TabsPrimitive.Trigger
	class={cn(
		'ring-offset-background focus-visible:ring-ring data-[state=active]:bg-background data-[state=active]:text-foreground inline-flex items-center justify-center rounded-sm px-3 py-1.5 text-sm font-medium whitespace-nowrap transition-all focus-visible:ring-2 focus-visible:ring-offset-2 focus-visible:outline-hidden disabled:pointer-events-none disabled:opacity-50 data-[state=active]:shadow-2xs',
		className
	)}
	{value}
	{...$$restProps}
	on:click
>
	<slot />
</TabsPrimitive.Trigger>
