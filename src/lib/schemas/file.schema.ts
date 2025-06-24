import { m } from '$lib/paraglide/messages';
import { z } from 'zod';

export const fileSchema = z.object({
	id: z.string(),
	name: z.string().min(1, m.file_name_is_required()),
	size: z.number().min(1, m.file_size_must_be_positive()),
	type: z.string().min(1, m.file_type_is_required()),
	file: z.string(),
	versions: z.array(z.string()),
	created: z.string(),
	updated: z.string()
});

export const fileCreateSchema = fileSchema.omit({ id: true, created: true, updated: true });
export const fileUpdateSchema = fileSchema.omit({ id: true, created: true, updated: true });
export const fileDeleteSchema = fileSchema.pick({ id: true });

export const fileUploadSchema = z.object({
	file: z
		.instanceof(File, { message: m.no_file_selected() })
		.refine((file) => file.size > 0, m.no_file_selected())
		.refine((file) => file.size <= 500 * 1024 * 1024, m.file_size_too_large()),
	versions: z.array(z.string()).min(1, m.version_required())
});

export type FileSchema = z.infer<typeof fileSchema>;
export type FileCreateSchema = z.infer<typeof fileCreateSchema>;
export type FileUpdateSchema = z.infer<typeof fileUpdateSchema>;
export type FileDeleteSchema = z.infer<typeof fileDeleteSchema>;
export type FileUploadSchema = z.infer<typeof fileUploadSchema>;
