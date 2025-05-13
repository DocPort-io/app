import { m } from '$lib/paraglide/messages';
import { z } from 'zod';

export const projectSchema = z.object({
	id: z.string(),
	name: z
		.string()
		.min(1, m.project_name_is_required())
		.max(100, m.project_name_cannot_be_longer_than_100_characters()),
	status: z.enum(['planned', 'active', 'completed']).default('active'),
	team: z.string(),
	created: z.string(),
	updated: z.string()
});

export const projectCreateSchema = projectSchema.omit({ id: true, created: true, updated: true });
export const projectUpdateSchema = projectSchema.omit({ id: true, created: true, updated: true });
export const projectDeleteSchema = projectSchema.pick({ id: true });

export type ProjectSchema = z.infer<typeof projectSchema>;
export type ProjectCreateSchema = z.infer<typeof projectCreateSchema>;
export type ProjectUpdateSchema = z.infer<typeof projectUpdateSchema>;
export type ProjectDeleteSchema = z.infer<typeof projectDeleteSchema>;
