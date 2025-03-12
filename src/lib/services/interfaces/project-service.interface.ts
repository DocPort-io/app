import type {
	ProjectCreateSchema,
	ProjectSchema,
	ProjectUpdateSchema
} from '$lib/schemas/project.schema';

export interface IProjectService {
	getProjects(): Promise<ProjectSchema[]>;
	createProject(data: ProjectCreateSchema): Promise<ProjectSchema>;
	updateProject(id: string, data: ProjectUpdateSchema): Promise<ProjectSchema>;
	deleteProject(id: string): Promise<void>;
}
