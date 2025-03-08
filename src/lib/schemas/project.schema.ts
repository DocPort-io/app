import { z } from 'zod';

export const projectSchema = z.object({
	name: z
		.string()
		.min(1, 'Project name is required')
		.max(255, 'Project name cannot be longer than 255 characters')
});

export type ProjectData = z.infer<typeof projectSchema>;
