import { z } from 'zod';

export const userSchema = z.object({
	id: z.string(),
	email: z.string(),
	emailVisibility: z.boolean(),
	verified: z.boolean(),
	name: z.string(),
	avatar: z.string(),
	created: z.string(),
	updated: z.string()
});

export type UserSchema = z.infer<typeof userSchema>;
