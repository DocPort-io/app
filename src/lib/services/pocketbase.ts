import type { ProjectSchema } from '$lib/schemas/project.schema';

import PocketBase, { RecordService } from 'pocketbase';

export interface TypedPocketBase extends PocketBase {
	collection(idOrName: string): RecordService; // default fallback for any other collection
	collection(idOrName: 'projects'): RecordService<ProjectSchema>;
}

export const getPocketBase = (): TypedPocketBase => {
	return new PocketBase('http://127.0.0.1:8090');
};
