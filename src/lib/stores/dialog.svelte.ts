export class DialogController<T> {
	isOpen = $state<boolean>(false);
	data = $state<T | undefined>(undefined);

	constructor(data?: T) {
		this.data = data;
	}

	open() {
		this.isOpen = true;
	}

	close() {
		this.isOpen = false;
	}
}

export const createDialogController = <T>(data?: T) => new DialogController<T>(data);
