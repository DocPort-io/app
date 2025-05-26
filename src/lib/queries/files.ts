import type { FileCreateSchema, FileUpdateSchema } from '$lib/schemas/file.schema';

import {
	createMutation,
	keepPreviousData,
	queryOptions,
	useQueryClient
} from '@tanstack/svelte-query';
import { FileService, type IFileService } from '$lib/services/file.service';
import { toast } from 'svelte-sonner';

const QUERY_BASE_KEY = 'files';

export type PaginatedFilesOptions = {
	version?: string;
	page: number;
	perPage: number;
	fileService?: IFileService;
};

export const createPaginatedFilesQuery = ({
	version,
	page,
	perPage,
	fileService
}: PaginatedFilesOptions) => {
	if (!fileService) fileService = new FileService();

	return queryOptions({
		queryKey: [QUERY_BASE_KEY, version, perPage, page],
		queryFn: () => fileService.findAll({ version: version!, page, perPage }),
		placeholderData: keepPreviousData,
		enabled: !!version
	});
};

export type FileMutationOptions = {
	fileService?: IFileService;
};

export const createAddFileMutation = ({ fileService }: FileMutationOptions = {}) => {
	if (!fileService) fileService = new FileService();
	const client = useQueryClient();

	return createMutation({
		mutationFn: (file: FileCreateSchema) => fileService.create(file),
		onSuccess: () => {
			client.invalidateQueries({ queryKey: [QUERY_BASE_KEY] });
			toast.success('File created successfully!');
		}
	});
};

export const createUpdateFileMutation = ({ fileService }: FileMutationOptions = {}) => {
	if (!fileService) fileService = new FileService();
	const client = useQueryClient();

	return createMutation({
		mutationFn: ({ id, file }: { id: string; file: FileUpdateSchema }) =>
			fileService.update(id, file),
		onSuccess: () => {
			client.invalidateQueries({ queryKey: [QUERY_BASE_KEY] });
			toast.success('File updated successfully!');
		}
	});
};

export const createDeleteFileMutation = ({ fileService }: FileMutationOptions = {}) => {
	if (!fileService) fileService = new FileService();
	const client = useQueryClient();

	return createMutation({
		mutationFn: (id: string) => fileService.remove(id),
		onSuccess: () => {
			client.invalidateQueries({ queryKey: [QUERY_BASE_KEY] });
			toast.success('File deleted successfully!');
		}
	});
};
