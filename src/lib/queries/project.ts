import { keepPreviousData, queryOptions } from '@tanstack/svelte-query';
import { ProjectService, type IProjectService } from '$lib/services/project.service';

const QUERY_BASE_KEY = 'project';

export type ProjectOptions = {
	id: string;
	projectService?: IProjectService;
};

export const createProjectQuery = ({ id, projectService }: ProjectOptions) => {
	if (!projectService) projectService = new ProjectService();

	return queryOptions({
		queryKey: [QUERY_BASE_KEY, id],
		queryFn: () => projectService.findOne(id),
		placeholderData: keepPreviousData
	});
};
