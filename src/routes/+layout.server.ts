import type { LayoutServerLoad } from './$types';

export const ssr = false;

export const load = (async () => {
	return {};
}) satisfies LayoutServerLoad;
