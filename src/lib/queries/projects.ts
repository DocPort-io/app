import type { ProjectCreateSchema, ProjectUpdateSchema } from '$lib/schemas/project.schema';

import {
	createMutation,
	keepPreviousData,
	queryOptions,
	useQueryClient
} from '@tanstack/svelte-query';
import { ProjectService, type IProjectService } from '$lib/services/project.service';
import { toast } from 'svelte-sonner';

const QUERY_BASE_KEY = 'projects';

export type PaginatedProjectsOptions = {
	team: string;
	page: number;
	perPage: number;
	projectService?: IProjectService;
};

export const createPaginatedProjectsQuery = ({
	team,
	page,
	perPage,
	projectService
}: PaginatedProjectsOptions) => {
	if (!projectService) projectService = new ProjectService();

	return queryOptions({
		queryKey: [QUERY_BASE_KEY, perPage, page],
		queryFn: () => projectService.getProjects({ team, page, perPage }),
		placeholderData: keepPreviousData
	});
};

export type ProjectMutationOptions = {
	projectService?: IProjectService;
};

export const createAddProjectMutation = ({ projectService }: ProjectMutationOptions = {}) => {
	if (!projectService) projectService = new ProjectService();
	const client = useQueryClient();

	return createMutation({
		mutationFn: (project: ProjectCreateSchema) => projectService.createProject(project),
		onSuccess: () => {
			client.invalidateQueries({ queryKey: [QUERY_BASE_KEY] });
			toast.success('Project created successfully!');
		}
	});
};

export const createUpdateProjectMutation = ({ projectService }: ProjectMutationOptions = {}) => {
	if (!projectService) projectService = new ProjectService();
	const client = useQueryClient();

	return createMutation({
		mutationFn: ({ id, project }: { id: string; project: ProjectUpdateSchema }) =>
			projectService.updateProject(id, project),
		onSuccess: () => {
			client.invalidateQueries({ queryKey: [QUERY_BASE_KEY] });
			toast.success('Project updated successfully!');
		}
	});
};

export const createDeleteProjectMutation = ({ projectService }: ProjectMutationOptions = {}) => {
	if (!projectService) projectService = new ProjectService();
	const client = useQueryClient();

	return createMutation({
		mutationFn: (id: string) => projectService.deleteProject(id),
		onSuccess: () => {
			client.invalidateQueries({ queryKey: [QUERY_BASE_KEY] });
			toast.success('Project deleted successfully!');
		}
	});
};
