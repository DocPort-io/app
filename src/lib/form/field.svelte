<script
	lang="ts"
	generics="Schema extends ZodObject<ZodRawShape>, FieldName extends keyof z.infer<Schema>"
>
	import type { Snippet } from 'svelte';
	import type z from 'zod';

	import { ZodObject, type ZodRawShape } from 'zod';

	import type { Form, FormField } from './form.svelte';

	type FieldProps = FormField<z.infer<Schema>[FieldName]> & {
		name: string;
	};

	type Props = {
		form: Form<Schema>;
		name: FieldName;
		children?: Snippet<[FieldProps]>;
	};

	const { form, name, children }: Props = $props();

	const formField = form.fields[name];

	const fieldProps: FieldProps = {
		name: String(name),
		...formField
	};
</script>

<div class="space-y-2">
	{@render children?.(fieldProps)}
</div>
