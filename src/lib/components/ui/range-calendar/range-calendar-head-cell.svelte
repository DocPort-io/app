<script lang="ts">
	import { cn } from '$lib/utils.js';
	import { RangeCalendar as RangeCalendarPrimitive } from 'bits-ui';

	type $$Props = RangeCalendarPrimitive.HeadCellProps;

	let className: $$Props['class'] = undefined;
	export { className as class };
</script>

<RangeCalendarPrimitive.HeadCell
	class={cn('text-muted-foreground w-9 rounded-md text-[0.8rem] font-normal', className)}
	{...$$restProps}
>
	<slot />
</RangeCalendarPrimitive.HeadCell>
