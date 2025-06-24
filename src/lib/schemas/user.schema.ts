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
export const userCreateSchema = userSchema.omit({ id: true, created: true, updated: true });
export const userUpdateSchema = userSchema.omit({ id: true, created: true, updated: true });
export const userDeleteSchema = userSchema.pick({ id: true });

export type UserSchema = z.infer<typeof userSchema>;
export type UserCreateSchema = z.infer<typeof userCreateSchema>;
export type UserUpdateSchema = z.infer<typeof userUpdateSchema>;
export type UserDeleteSchema = z.infer<typeof userDeleteSchema>;
