<script lang="ts">
	import { cn } from '$lib/utils.js';
	import { RangeCalendar as RangeCalendarPrimitive } from 'bits-ui';

	type $$Props = RangeCalendarPrimitive.HeaderProps;

	let className: $$Props['class'] = undefined;
	export { className as class };
</script>

<RangeCalendarPrimitive.Header
	class={cn('relative flex w-full items-center justify-between pt-1', className)}
	{...$$restProps}
>
	<slot />
</RangeCalendarPrimitive.Header>
