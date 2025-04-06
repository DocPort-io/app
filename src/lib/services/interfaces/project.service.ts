import type {
	ProjectCreateSchema,
	ProjectSchema,
	ProjectUpdateSchema
} from '$lib/schemas/project.schema';

export type ProjectServiceGetProjectsOptions = {
	page?: number;
	perPage?: number;
	team: string;
};

export type ProjectServiceGetProjectsResult = {
	items: ProjectSchema[];
	totalItems: number;
	totalPages: number;
};

export interface IProjectService {
	getProjects(options: ProjectServiceGetProjectsOptions): Promise<ProjectServiceGetProjectsResult>;
	createProject(data: ProjectCreateSchema): Promise<ProjectSchema>;
	updateProject(id: string, data: ProjectUpdateSchema): Promise<ProjectSchema>;
	deleteProject(id: string): Promise<void>;
}
