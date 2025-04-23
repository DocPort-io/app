export interface IPaginationController {
	page: number;
	perPage: number;
}

export type PaginationControllerOptions = {
	page?: number;
	perPage?: number;
};

export class PaginationController implements IPaginationController {
	page = $state<number>(1);
	perPage = $state<number>(20);
	totalItems = $state<number>(0);
	totalPages = $state<number>(0);

	constructor({ page, perPage }: PaginationControllerOptions = {}) {
		if (page) this.page = page;
		if (perPage) this.perPage = perPage;
	}
}

export const createPaginationController = (options?: PaginationControllerOptions) =>
	new PaginationController(options);
