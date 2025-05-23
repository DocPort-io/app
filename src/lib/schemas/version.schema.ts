import { m } from '$lib/paraglide/messages';
import { z } from 'zod';

export const versionSchema = z.object({
	id: z.string(),
	name: z
		.string()
		.min(1, m.version_name_is_required())
		.max(100, m.version_name_cannot_be_longer_than_100_characters()),
	description: z.string().max(300, m.description_cannot_be_longer_than_300_characters()).optional(),
	project: z.string(),
	files: z.array(z.any()),
	created: z.string(),
	updated: z.string()
});

export const versionCreateSchema = versionSchema.omit({ id: true, created: true, updated: true });
export const versionUpdateSchema = versionSchema.omit({ id: true, created: true, updated: true });
export const versionDeleteSchema = versionSchema.pick({ id: true });

export type VersionSchema = z.infer<typeof versionSchema>;
export type VersionCreateSchema = z.infer<typeof versionCreateSchema>;
export type VersionUpdateSchema = z.infer<typeof versionUpdateSchema>;
export type VersionDeleteSchema = z.infer<typeof versionDeleteSchema>;
