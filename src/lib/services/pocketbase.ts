import type { ProjectSchema } from '$lib/schemas/project.schema';
import type { UserSchema } from '$lib/schemas/user.schema';

import PocketBase, { LocalAuthStore, RecordService } from 'pocketbase';

export interface TypedPocketBase extends PocketBase {
	collection(idOrName: string): RecordService; // default fallback for any other collection
	collection(idOrName: 'projects'): RecordService<ProjectSchema>;
	collection(idOrName: 'users'): RecordService<UserSchema>;
}

const authStore = new LocalAuthStore();
const pocketBase = new PocketBase('http://localhost:8080', authStore) as TypedPocketBase;
// const pocketBase = new PocketBase('https://api.docport.io', authStore) as TypedPocketBase;

export const getPocketBase = () => pocketBase;
