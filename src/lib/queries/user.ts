import { keepPreviousData, queryOptions } from '@tanstack/svelte-query';
import { UserService, type IUserService } from '$lib/services/user.service';

const QUERY_BASE_KEY = 'user';

export type UserOptions = {
	id: string;
	userService?: IUserService;
};

export const createUserQuery = ({ id, userService }: UserOptions) => {
	if (!userService) userService = new UserService();

	return queryOptions({
		queryKey: [QUERY_BASE_KEY, id],
		queryFn: () => userService.findOne(id),
		placeholderData: keepPreviousData
	});
};
