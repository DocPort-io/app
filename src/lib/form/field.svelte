<script
	lang="ts"
	generics="TSchema extends ZodObject<ZodRawShape>, TFieldName extends keyof z.infer<TSchema>"
>
	import type { ClassValue } from 'svelte/elements';
	import type { z } from 'zod';

	import { cn } from '$lib/utils';
	import { type Snippet } from 'svelte';
	import { ZodObject, type ZodRawShape } from 'zod';

	import { setFormField } from './field.context.svelte';
	import { type Form, type FormField } from './form.svelte';

	/**
	 * Props for the Field component
	 */
	type Props = {
		/** Reference to the field container element */
		ref?: HTMLElement | null;
		/** Additional CSS classes */
		class?: ClassValue | null;
		/** The form instance this field belongs to */
		form: Form<TSchema>;
		/** The name of the field in the schema */
		name: TFieldName;
		/** Render prop that provides field props and state */
		children?: Snippet<
			[
				{
					/** Props to spread on the input element */
					props: FormField<z.infer<TSchema>[TFieldName]>['props'];
					/** Current field state for conditional rendering */
					state: FormField<z.infer<TSchema>[TFieldName]>['state'];
				}
			]
		>;
	};

	let {
		ref = $bindable(null),
		class: className,
		form,
		name,
		children,
		...restProps
	}: Props = $props();

	// Get the field instance for this specific field name
	const formField = form.fields[name];

	// Provide field context for child components (like FieldErrors, FieldLabel)
	setFormField(formField);
</script>

<!--
	Field container with consistent spacing and styling
	Provides context for field-related components
-->
<div bind:this={ref} class={cn('space-y-2', className)} {...restProps}>
	{@render children?.({ props: formField.props, state: formField.state })}
</div>
