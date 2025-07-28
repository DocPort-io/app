<script lang="ts">
	import type { HTMLAttributes } from 'svelte/elements';

	import { cn } from '$lib/utils';

	type Props = HTMLAttributes<HTMLParagraphElement> & {
		ref?: HTMLParagraphElement | null;
	};

	let { ref = $bindable(null), children, class: className, ...restProps }: Props = $props();
</script>

<p
	bind:this={ref}
	class={cn('text-muted-foreground text-sm', className)}
	data-slot="form-description"
	{...restProps}
>
	{@render children?.()}
</p>
