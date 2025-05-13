import { m } from '$lib/paraglide/messages';
import { z } from 'zod';

export const signInSchema = z.object({
	email: z.string().email(m.that_email_address_does_not_look_right()),
	password: z.string()
});
