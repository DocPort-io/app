import type z from 'zod';

import * as uuid from 'uuid';
import { ZodError, type ZodObject, type ZodRawShape } from 'zod';

export type OnSubmitOptions<Schema extends ZodObject<ZodRawShape>> = {
	data: z.infer<Schema>;
	state: FormState;
};

export type CreateFormOptions<Schema extends ZodObject<ZodRawShape>> = {
	schema: Schema;
	defaultValues?: Partial<z.infer<Schema>>;
	onSubmit?: (options: OnSubmitOptions<Schema>) => Promise<void> | void;
};

export type FormState = {
	isSubmittable: boolean;
	isSubmitting: boolean;
	isValidating: boolean;
	isValid: boolean;
};

export type Form<Schema extends ZodObject<ZodRawShape>> = {
	state: FormState;
	fields: {
		[FieldName in keyof z.infer<Schema>]: FormField<z.infer<Schema>[FieldName]>;
	};
	props: {
		onsubmit: (event: SubmitEvent) => void;
	};
};

export const createForm = <Schema extends ZodObject<ZodRawShape>>(
	options: CreateFormOptions<Schema>
): Form<Schema> => {
	const { schema, defaultValues, onSubmit } = options;

	type SchemaType = z.infer<Schema>;

	let isSubmitting = $state(false);
	let isValidating = $state(false);
	let isValid = $state(true);
	const isSubmittable = $derived(isValid && !isValidating && !isSubmitting);

	const state = () => ({
		isSubmittable,
		isSubmitting,
		isValidating,
		isValid
	});

	// Initialize fields based on schema shape
	const schemaShape = schema.shape;

	const fields = $state(
		Object.fromEntries(
			Object.keys(schemaShape).map((key) => {
				const typedKey = key as keyof SchemaType;
				const defaultValue = defaultValues?.[typedKey];
				return [key, createFormField({ defaultValue, name: key })];
			})
		) as { [FieldName in keyof SchemaType]: FormField<SchemaType[FieldName]> }
	);

	const getFieldValues = () =>
		Object.fromEntries(
			Object.entries(fields).map(([key, field]) => {
				return [key, field.state.value];
			})
		);

	const handleSubmit = async (event: SubmitEvent) => {
		event.preventDefault();
		event.stopPropagation();

		if (!isSubmittable) return;

		isSubmitting = true;

		const getData = (): SchemaType => {
			const fieldValues = getFieldValues();

			try {
				return schema.parse(fieldValues);
			} catch {
				return {} as SchemaType;
			}
		};

		try {
			await onSubmit?.({
				data: getData(),
				state: state()
			});
		} catch (error) {
			console.error('Error during form submission:', error);
		}

		isSubmitting = false;
	};

	const validateFields = () => {
		console.log('Validating fields...');

		isValidating = true;

		try {
			const fieldValues = getFieldValues();
			schema.parse(fieldValues);

			isValid = true;
		} catch (error) {
			isValid = false;
			console.error('Validation failed:', error);

			if (error instanceof ZodError) {
				console.log(error.issues);
			}
		}

		isValidating = false;
	};

	$effect(() => {
		validateFields();
	});

	return {
		get state() {
			return state();
		},
		fields,
		props: {
			onsubmit: handleSubmit
		}
	};
};

type CreateFormFieldOptions<FieldType> = {
	defaultValue?: FieldType;
	name: string;
};

type FormFieldState<FieldType> = {
	value: FieldType | undefined;
};

export type FormField<FieldType> = {
	state: FormFieldState<FieldType>;
	props: {
		id: string;
		name: string;
		onchange: () => void;
		onblur: () => void;
	};
};

export const createFormField = <FieldType>(
	options: CreateFormFieldOptions<FieldType>
): FormField<FieldType> => {
	const { defaultValue, name } = options;

	const state = $state<FormFieldState<FieldType>>({
		value: defaultValue
	});

	const handleChange = () => {
		// Handle change logic
	};

	const handleBlur = () => {
		// Handle blur logic
	};

	const id = `form-field-${name}-${uuid.v4()}`;

	return {
		state,
		props: {
			id,
			name,
			onchange: handleChange,
			onblur: handleBlur
		}
	};
};
