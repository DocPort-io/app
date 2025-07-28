<script
	lang="ts"
	generics="TSchema extends ZodObject<ZodRawShape>, TFieldName extends keyof z.infer<TSchema>"
>
	import type { HTMLAttributes } from 'svelte/elements';
	import type { z } from 'zod';

	import { cn } from '$lib/utils';
	import { ZodObject, type ZodRawShape } from 'zod';

	import { getFormField } from './field.context.svelte';

	type Props = HTMLAttributes<HTMLDivElement> & {
		ref?: HTMLDivElement | null;
	};

	let { ref = $bindable(null), class: className, ...restProps }: Props = $props();

	const field = getFormField<TSchema, TFieldName>();
</script>

{#if field.state.errors.length > 0}
	<div bind:this={ref} class={cn('space-y-1', className)} {...restProps}>
		{#each field.state.errors as error, index (index)}
			<p class="text-destructive text-sm font-medium">{error}</p>
		{/each}
	</div>
{/if}
