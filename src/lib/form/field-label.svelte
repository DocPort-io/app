<script
	lang="ts"
	generics="Schema extends ZodObject<ZodRawShape>, FieldName extends keyof z.infer<Schema>"
>
	import type z from 'zod';

	import { Label } from '$lib/components/ui/label';
	import { cn } from '$lib/utils';
	import { Label as LabelPrimitive } from 'bits-ui';
	import { ZodObject, type ZodRawShape } from 'zod';

	import { getFormField } from './field.context.svelte';

	let {
		ref = $bindable(null),
		children,
		class: className,
		...restProps
	}: LabelPrimitive.RootProps = $props();

	const fieldProps = getFormField<Schema, FieldName>();
</script>

<Label
	bind:ref
	class={cn('data-[fs-error]:text-destructive', className)}
	data-slot="form-label"
	for={fieldProps.props.id}
	{...restProps}
>
	{@render children?.()}
</Label>
