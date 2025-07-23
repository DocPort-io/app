<script
	lang="ts"
	generics="Schema extends ZodObject<ZodRawShape>, FieldName extends keyof z.infer<Schema>"
>
	import type { ClassValue } from 'svelte/elements';
	import type z from 'zod';

	import { cn } from '$lib/utils';
	import { type Snippet } from 'svelte';
	import { ZodObject, type ZodRawShape } from 'zod';

	import { setFormField } from './field.context.svelte';
	import { type Form, type FormField } from './form.svelte';

	type Props = {
		ref?: HTMLElement | null;
		class?: ClassValue | null;
		form: Form<Schema>;
		name: FieldName;
		children?: Snippet<[FormField<z.infer<Schema>[FieldName]>]>;
	};

	let {
		ref = $bindable(null),
		class: className,
		form,
		name,
		children,
		...restProps
	}: Props = $props();

	const formField = form.fields[name];

	setFormField(formField);
</script>

<div bind:this={ref} class={cn('space-y-2', className)} {...restProps}>
	{@render children?.(formField)}
</div>
