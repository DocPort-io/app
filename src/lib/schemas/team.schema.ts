import { z } from 'zod';

export const teamSchema = z.object({
	id: z.string(),
	name: z
		.string()
		.min(1, 'Team name is required')
		.max(100, 'Team name cannot be longer than 100 characters'),
	user: z.array(z.string()),
	created: z.string(),
	updated: z.string()
});

export const teamCreateSchema = teamSchema.omit({ id: true, created: true, updated: true });
export const teamUpdateSchema = teamSchema.omit({ id: true, created: true, updated: true });
export const teamDeleteSchema = teamSchema.pick({ id: true });

export type TeamSchema = z.infer<typeof teamSchema>;
export type TeamCreateSchema = z.infer<typeof teamCreateSchema>;
export type TeamUpdateSchema = z.infer<typeof teamUpdateSchema>;
export type TeamDeleteSchema = z.infer<typeof teamDeleteSchema>;
