import { keepPreviousData, queryOptions } from '@tanstack/svelte-query';
import { TeamService, type ITeamService } from '$lib/services/team.service';

const QUERY_BASE_KEY = 'teams';

export type PaginatedProjectsOptions = {
	page?: number;
	perPage?: number;
	teamService?: ITeamService;
};

export const createPaginatedTeamsQuery = ({
	page,
	perPage,
	teamService
}: PaginatedProjectsOptions = {}) => {
	if (!teamService) teamService = new TeamService();

	page = page ?? 1;
	perPage = perPage ?? 500;

	return queryOptions({
		queryKey: [QUERY_BASE_KEY, perPage, page],
		queryFn: () => teamService.findAll({ page, perPage }),
		placeholderData: keepPreviousData
	});
};
