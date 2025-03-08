import PocketBase, { type RecordModel } from 'pocketbase';
import { getContext, setContext } from 'svelte';

type Project = RecordModel & {
	name: string;
	created: string;
	updated: string;
};

export class ProjectsState {
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

	async add() {
		const projectData: Partial<Project> = {
			name: `Project ${crypto.randomUUID()}`
		};

		await this.#pocketBase.collection<Project>('projects').create(projectData);
		this.load();
	}
}

const PROJECTS_STATE_KEY = Symbol('PROJECTS_STATE');

export const setProjectsState = () => {
	return setContext(PROJECTS_STATE_KEY, new ProjectsState());
};

export const getProjectsState = () => {
	return getContext<ReturnType<typeof setProjectsState>>(PROJECTS_STATE_KEY);
};
