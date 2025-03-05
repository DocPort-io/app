import { getContext, setContext } from 'svelte';

export class UserState {
	name = $state<string>('Jonas');
}

const USER_STATE_KEY = Symbol('USER_STATE');

export const setUserState = () => {
	return setContext(USER_STATE_KEY, new UserState());
};

export const getUserState = () => {
	return getContext<ReturnType<typeof setUserState>>(USER_STATE_KEY);
};
