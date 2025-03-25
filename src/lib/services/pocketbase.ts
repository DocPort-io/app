import type { ProjectSchema } from '$lib/schemas/project.schema';
import type { UserSchema } from '$lib/schemas/user.schema';

import { env } from '$env/dynamic/public';
import PocketBase, { LocalAuthStore, type RecordService } from 'pocketbase';

export interface TypedPocketBase extends PocketBase {
	collection(idOrName: string): RecordService; // default fallback for any other collection
	collection(idOrName: 'projects'): RecordService<ProjectSchema>;
	collection(idOrName: 'users'): RecordService<UserSchema>;
}

const authStore = new LocalAuthStore();
const pocketBase = new PocketBase(env.PUBLIC_POCKETBASE_URL, authStore) as TypedPocketBase;

export const getPocketBase = () => pocketBase;
