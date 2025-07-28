<script
	lang="ts"
	generics="TSchema extends ZodObject<ZodRawShape>, TFieldName extends keyof z.infer<TSchema>"
>
	import type { z } from 'zod';

	import { Label } from '$lib/components/ui/label';
	import { cn } from '$lib/utils';
	import { Label as LabelPrimitive } from 'bits-ui';
	import { ZodObject, type ZodRawShape } from 'zod';

	import { getFormField } from './field.context.svelte';

	/**
	 * Props for the FieldLabel component, extending the base Label props
	 */
	let {
		ref = $bindable(null),
		children,
		class: className,
		...restProps
	}: LabelPrimitive.RootProps = $props();

	// Get the field context to access field state and ID
	const field = getFormField<TSchema, TFieldName>();
</script>

<!--
	Accessible label component that automatically:
	- Associates with the field input via ID
	- Changes color when field has errors
	- Maintains consistent styling and accessibility
-->
<Label
	bind:ref
	class={cn(
		'text-sm leading-none font-medium peer-disabled:cursor-not-allowed peer-disabled:opacity-70',
		field.state.errors.length > 0 && 'text-destructive',
		className
	)}
	data-slot="form-label"
	for={field.props.id}
	{...restProps}
>
	{@render children?.()}
</Label>
