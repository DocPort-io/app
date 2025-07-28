<script lang="ts" generics="TSchema extends ZodObject<ZodRawShape>">
	import type { HTMLAttributes } from 'svelte/elements';

	import { AlertTriangle } from '@lucide/svelte';
	import * as Alert from '$lib/components/ui/alert';
	import { cn } from '$lib/utils';
	import { ZodObject, type ZodRawShape } from 'zod';

	import type { Form } from './form.svelte';

	type Props = HTMLAttributes<HTMLDivElement> & {
		form: Form<TSchema>;
		ref?: HTMLDivElement | null;
	};

	let { form, ref = $bindable(null), class: className, ...restProps }: Props = $props();

	const submitErrors = $derived(form.state.errors._submit || []);
</script>

{#if submitErrors.length > 0}
	<Alert.Root bind:ref class={cn(className)} variant="destructive" {...restProps}>
		<AlertTriangle />
		<Alert.Description>
			{#each submitErrors as error, index (index)}
				<p class={cn('font-medium', { 'mt-2': index > 0 })}>{error}</p>
			{/each}
		</Alert.Description>
	</Alert.Root>
{/if}
