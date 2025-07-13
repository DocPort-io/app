import type { ZodObject, ZodRawShape } from 'zod';
import type z from 'zod';

export type OnSubmitOptions<Schema extends ZodObject<ZodRawShape>> = {
	data: z.infer<Schema>;
};

export type CreateFormOptions<Schema extends ZodObject<ZodRawShape>> = {
	schema: Schema;
	defaultValues?: Partial<z.infer<Schema>>;
	onSubmit: (options: OnSubmitOptions<Schema>) => Promise<void> | void;
};

export type FormState = {
	isSubmittable: boolean;
	isSubmitting: boolean;
	isSubmitted: boolean;
};

export type Form<Schema extends ZodObject<ZodRawShape>> = {
	handleSubmit: () => void;
	state: FormState;
	fields: Record<keyof z.infer<Schema>, FormField>;
};

export const createForm = <Schema extends ZodObject<ZodRawShape>>(
	options: CreateFormOptions<Schema>
): Form<Schema> => {
	const { schema, defaultValues, onSubmit } = options;

	type SchemaType = z.infer<Schema>;
	type FieldName = keyof SchemaType;

	const state = $state<FormState>({
		isSubmittable: false,
		isSubmitting: false,
		isSubmitted: false
	});

	// Initialize fields based on schema shape
	const schemaShape = schema.shape;
	const initialFields: Record<string, FormField> = {};

	for (const key in schemaShape) {
		const defaultValue = defaultValues?.[key];
		initialFields[key] = createFormField({ defaultValue });
	}

	const fields = $state(initialFields as Record<FieldName, FormField>);

	const handleSubmit = async () => {
		state.isSubmitting = true;
		state.isSubmitted = false;

		const fieldValues = Object.fromEntries(
			Object.entries(fields).map(([key, field]) => {
				return [key, field.state.value];
			})
		);

		await onSubmit({
			data: schema.parse(fieldValues)
		});

		state.isSubmitting = false;
		state.isSubmitted = true;
	};

	return {
		handleSubmit,
		state,
		fields
	};
};

type CreateFormFieldOptions = {
	defaultValue?: string;
};

type FormFieldState = {
	value: string;
};

export type FormField = {
	state: FormFieldState;
	handleChange: (value: string) => void;
	handleBlur: () => void;
};

const createFormField = (options: CreateFormFieldOptions): FormField => {
	const { defaultValue = '' } = options;

	const state = $state({
		value: defaultValue
	});

	const handleChange = (_value: string) => {
		state.value = _value;
	};

	const handleBlur = () => {
		// Handle blur logic if needed
	};

	return {
		state,
		handleChange,
		handleBlur
	};
};
