import type { VersionCreateSchema, VersionUpdateSchema } from '$lib/schemas/version.schema';

import {
	createMutation,
	keepPreviousData,
	queryOptions,
	useQueryClient
} from '@tanstack/svelte-query';
import { VersionService, type IVersionService } from '$lib/services/version.service';
import { toast } from 'svelte-sonner';

const QUERY_BASE_KEY = 'versions';

export type PaginatedVersionsOptions = {
	project?: string;
	page: number;
	perPage: number;
	versionService?: IVersionService;
};

export const createPaginatedVersionsQuery = ({
	project,
	page,
	perPage,
	versionService
}: PaginatedVersionsOptions) => {
	if (!versionService) versionService = new VersionService();

	return queryOptions({
		queryKey: [QUERY_BASE_KEY, project, perPage, page],
		queryFn: () => versionService.findAll({ project: project!, page, perPage }),
		placeholderData: keepPreviousData,
		enabled: !!project
	});
};

export type VersionMutationOptions = {
	versionService?: IVersionService;
};

export const createAddVersionMutation = ({ versionService }: VersionMutationOptions = {}) => {
	if (!versionService) versionService = new VersionService();
	const client = useQueryClient();

	return createMutation({
		mutationFn: (version: VersionCreateSchema) => versionService.create(version),
		onSuccess: () => {
			client.invalidateQueries({ queryKey: [QUERY_BASE_KEY] });
			toast.success('Version created successfully!');
		}
	});
};

export const createUpdateVersionMutation = ({ versionService }: VersionMutationOptions = {}) => {
	if (!versionService) versionService = new VersionService();
	const client = useQueryClient();

	return createMutation({
		mutationFn: ({ id, version }: { id: string; version: VersionUpdateSchema }) =>
			versionService.update(id, version),
		onSuccess: () => {
			client.invalidateQueries({ queryKey: [QUERY_BASE_KEY] });
			toast.success('Version updated successfully!');
		}
	});
};

export const createDeleteVersionMutation = ({ versionService }: VersionMutationOptions = {}) => {
	if (!versionService) versionService = new VersionService();
	const client = useQueryClient();

	return createMutation({
		mutationFn: (id: string) => versionService.remove(id),
		onSuccess: () => {
			client.invalidateQueries({ queryKey: [QUERY_BASE_KEY] });
			toast.success('Version deleted successfully!');
		}
	});
};
