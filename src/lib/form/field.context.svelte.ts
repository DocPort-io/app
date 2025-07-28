import type { z, ZodObject, ZodRawShape } from 'zod';

import { getContext, setContext } from 'svelte';

import type { FormField } from './form.svelte';

const FORM_FIELD_KEY = Symbol('FORM_FIELD_PROPS');

export const setFormField = <
	TSchema extends ZodObject<ZodRawShape>,
	TFieldName extends keyof z.infer<TSchema>
>(
	formField: FormField<z.infer<TSchema>[TFieldName]>
) => {
	return setContext(FORM_FIELD_KEY, formField);
};

export const getFormField = <
	TSchema extends ZodObject<ZodRawShape>,
	TFieldName extends keyof z.infer<TSchema>
>() => {
	return getContext<FormField<z.infer<TSchema>[TFieldName]>>(FORM_FIELD_KEY);
};
