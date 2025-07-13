<script lang="ts" generics="T extends ZodObject<ZodRawShape>">
	import type { Snippet } from 'svelte';

	import { ZodObject, type ZodRawShape } from 'zod';

	import type { Form, FormField } from './form.svelte';

	type FieldProps = FormField & {
		name: string;
	};

	type Props = {
		form: Form<T>;
		name: keyof Form<T>['fields'];
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
