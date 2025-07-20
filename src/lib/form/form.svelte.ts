import type { ZodObject, ZodRawShape } from 'zod';
import type z from 'zod';

export type OnSubmitOptions<Schema extends ZodObject<ZodRawShape>> = {
	data: z.infer<Schema>;
};

export type CreateFormOptions<Schema extends ZodObject<ZodRawShape>> = {
	schema: Schema;
	defaultValues?: Partial<z.infer<Schema>>;
	onSubmit?: (options: OnSubmitOptions<Schema>) => Promise<void> | void;
};

export type FormState = {
	isSubmittable: boolean;
	isSubmitting: boolean;
	isSubmitted: boolean;
};

export type Form<Schema extends ZodObject<ZodRawShape>> = {
	handleSubmit: () => void;
	state: FormState;
	fields: {
		[FieldName in keyof z.infer<Schema>]: FormField<z.infer<Schema>[FieldName]>;
	};
};

export const createForm = <Schema extends ZodObject<ZodRawShape>>(
	options: CreateFormOptions<Schema>
): Form<Schema> => {
	const { schema, defaultValues, onSubmit } = options;

	type SchemaType = z.infer<Schema>;

	const state = $state<FormState>({
		isSubmittable: false,
		isSubmitting: false,
		isSubmitted: false
	});

	// Initialize fields based on schema shape
	const schemaShape = schema.shape;

	const fieldsEntries = Object.keys(schemaShape).map((key) => {
		const typedKey = key as keyof SchemaType;
		const defaultValue = defaultValues?.[typedKey];
		return [key, createFormField({ defaultValue })];
	});

	const initialFields = Object.fromEntries(fieldsEntries) as {
		[FieldName in keyof SchemaType]: FormField<SchemaType[FieldName]>;
	};

	const fields = $state(initialFields);

	const handleSubmit = async () => {
		state.isSubmitting = true;
		state.isSubmitted = false;

		const fieldValues = Object.fromEntries(
			Object.entries(fields).map(([key, field]) => {
				return [key, field.state.value];
			})
		);

		await onSubmit?.({
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

type CreateFormFieldOptions<FieldType> = {
	defaultValue?: FieldType;
};

type FormFieldState<FieldType> = {
	value: FieldType | undefined;
};

export type FormField<FieldType> = {
	state: FormFieldState<FieldType>;
	handleChange: (value: FieldType) => void;
	handleBlur: () => void;
};

const createFormField = <FieldType>(
	options: CreateFormFieldOptions<FieldType>
): FormField<FieldType> => {
	const { defaultValue } = options;

	const state = $state<FormFieldState<FieldType>>({
		value: defaultValue
	});

	const handleChange = (_value: FieldType) => {
		state.value = _value;
	};

	const handleBlur = () => {
		// Handle blur logic
	};

	return {
		state,
		handleChange,
		handleBlur
	};
};
