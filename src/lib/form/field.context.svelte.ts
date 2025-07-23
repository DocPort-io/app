import type { z, ZodObject, ZodRawShape } from 'zod';

import { getContext, setContext } from 'svelte';

import type { FormField } from './form.svelte';

const FORM_FIELD_KEY = Symbol('FORM_FIELD_PROPS');

export const setFormField = <
	Schema extends ZodObject<ZodRawShape>,
	FieldName extends keyof z.infer<Schema>
>(
	formField: {
		[FieldName in keyof z.infer<Schema>]: FormField<z.infer<Schema>[FieldName]>;
	}[FieldName]
) => {
	return setContext(FORM_FIELD_KEY, formField);
};

export const getFormField = <
	Schema extends ZodObject<ZodRawShape>,
	FieldName extends keyof z.infer<Schema>
>() => {
	return getContext<ReturnType<typeof setFormField<Schema, FieldName>>>(FORM_FIELD_KEY);
};
