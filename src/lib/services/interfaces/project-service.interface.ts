import type {
	ProjectCreateSchema,
	ProjectSchema,
	ProjectUpdateSchema
} from '$lib/schemas/project.schema';

export interface IProjectService {
	getProjects(options: {
		page?: number;
		perPage?: number;
		filter?: string;
	}): Promise<{ items: ProjectSchema[]; totalItems: number; totalPages: number }>;
	createProject(data: ProjectCreateSchema): Promise<ProjectSchema>;
	updateProject(id: string, data: ProjectUpdateSchema): Promise<ProjectSchema>;
	deleteProject(id: string): Promise<void>;
}
