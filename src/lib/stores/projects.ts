import { asyncWritable } from '@square/svelte-store';
import { derived, writable } from 'svelte/store';
import PocketBase, { type RecordModel } from 'pocketbase';

const pb = new PocketBase('http://127.0.0.1:8090');

// Error handling store
export const error = writable<string | null>(null);

// Create an asyncWritable store for projects
export const projects = asyncWritable<[], Omit<RecordModel, 'collectionId' | 'collectionName'>[]>(
	[],
	async () => {
		const records = await pb.collection('projects').getList(1, 50, {
			sort: '-created'
		});
		return records.items;
	},
	undefined,
	{
		reloadable: true,
		trackState: true
	}
);

// Project management actions with optimistic updates
export const projectActions = {
	async createProject() {
		const projectData: Omit<RecordModel, 'id'> = {
			name: 'Projectz X'
		};

		// Generate a temporary ID for optimistic update
		const tempId = `temp_${Date.now()}_${Math.random().toString(36).substr(2, 9)}`;

		// Create new project with temp ID and optimistic flag
		const tempProject = {
			id: tempId,
			created: new Date().toISOString(),
			updated: new Date().toISOString(),
			...projectData,
			_isOptimistic: true
		};

		// Update the store optimistically
		projects.update((currentProjects) => [tempProject, ...currentProjects]);

		try {
			// Perform actual API call
			const record = await pb.collection('projects').create(projectData);

			// Update our list with the server-generated record
			projects.update((currentProjects) =>
				currentProjects.map((project) => (project.id === tempId ? { ...record } : project))
			);

			return record;
		} catch (err) {
			// Remove optimistic item on error
			projects.update((items) => items.filter((item) => item.id !== tempId));
			throw err;
		}
	}
};
