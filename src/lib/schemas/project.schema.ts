import { m } from '$lib/paraglide/messages';
import { z } from 'zod';

export const projectSchema = z.object({
	id: z.string(),
	name: z.string().min(1, m.tame_candid_platypus_care()).max(100, m.long_gross_stingray_snip()),
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
