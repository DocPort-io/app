import { m } from '$lib/paraglide/messages';
import { z } from 'zod';

export const signInSchema = z.object({
	email: z.string().email(m.broad_giant_impala_trim()),
	password: z.string()
});
