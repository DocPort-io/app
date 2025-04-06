export type PaginationControllerOptions = {
	page?: number;
	perPage?: number;
	totalItems?: number;
	totalPages?: number;
};

export class PaginationController {
	page = $state<number>(1);
	perPage = $state<number>(20);
	totalItems = $state<number>(0);
	totalPages = $state<number>(0);

	constructor({ page, perPage, totalItems, totalPages }: PaginationControllerOptions = {}) {
		if (page) this.page = page;
		if (perPage) this.perPage = perPage;
		if (totalItems) this.totalItems = totalItems;
		if (totalPages) this.totalPages = totalPages;
	}
}

export const createPaginationController = (options?: PaginationControllerOptions) =>
	new PaginationController(options);
