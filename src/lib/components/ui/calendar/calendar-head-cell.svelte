<script lang="ts">
	import { cn } from '$lib/utils.js';
	import { Calendar as CalendarPrimitive } from 'bits-ui';

	type $$Props = CalendarPrimitive.HeadCellProps;

	let className: $$Props['class'] = undefined;
	export { className as class };
</script>

<CalendarPrimitive.HeadCell
	class={cn('text-muted-foreground w-9 rounded-md text-[0.8rem] font-normal', className)}
	{...$$restProps}
>
	<slot />
</CalendarPrimitive.HeadCell>
