import type { ProjectCreateSchema, ProjectDeleteSchema } from '$lib/schemas/project.schema';
import PocketBase, { type RecordModel } from 'pocketbase';
import { getContext, setContext } from 'svelte';

type Project = RecordModel & {
	name: string;
	created: string;
	updated: string;
};

export class Projects {
	#pocketBase: PocketBase;
	projects = $state<Project[]>([]);
	loadingPromise = $state<Promise<void> | undefined>();

	constructor(pocketBase?: PocketBase) {
		this.#pocketBase = pocketBase ?? new PocketBase('http://127.0.0.1:8090');
	}

	load() {
		this.loadingPromise = this.#fetch();
	}

	#fetch = async () => {
		await new Promise((resolve) => setTimeout(resolve, 500));

		const records = await this.#pocketBase.collection<Project>('projects').getList(1, 50, {
			sort: '-created'
		});

		this.projects = records.items;
	};

	async add(project: ProjectCreateSchema) {
		await this.#pocketBase.collection<Project>('projects').create(project);
		await this.#fetch();
	}

	async remove(project: ProjectDeleteSchema) {
		await this.#pocketBase.collection<Project>('projects').delete(project.id);
		await this.#fetch();
	}
}

const PROJECTS_KEY = Symbol('PROJECTS');

export const setProjects = () => {
	return setContext(PROJECTS_KEY, new Projects());
};

export const getProjects = () => {
	return getContext<ReturnType<typeof setProjects>>(PROJECTS_KEY);
};
