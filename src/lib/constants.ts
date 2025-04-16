export const AppRoute = {
	DASHBOARD: () => '/dashboard',
	PROJECTS: () => '/projects',
	PROJECT_VIEW: (id: string) => `/projects/${id}`,
	LOGIN: () => '/auth/login'
};
