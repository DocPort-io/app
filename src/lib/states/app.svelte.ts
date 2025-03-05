import { getContext, setContext } from 'svelte';

type Theme = 'light' | 'dark' | 'system' | undefined;

export class AppState {
	theme = $state<Theme>();

	activateSystemTheme() {
		this.theme = 'system';
	}

	activateLightTheme() {
		this.theme = 'light';
	}

	activateDarkTheme() {
		this.theme = 'dark';
	}
}

const APP_STATE_KEY = Symbol('APP_STATE');

export const setAppState = () => {
	return setContext(APP_STATE_KEY, new AppState());
};

export const getAppState = () => {
	return getContext<ReturnType<typeof setAppState>>(APP_STATE_KEY);
};
