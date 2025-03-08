<script lang="ts">
	import { AppRoute } from '$lib/constants';
	import Paperclip from 'lucide-svelte/icons/paperclip';
	import Search from 'lucide-svelte/icons/search';
	import CircleUser from 'lucide-svelte/icons/circle-user';
	import Sun from 'lucide-svelte/icons/sun';
	import Moon from 'lucide-svelte/icons/moon';
	import Computer from 'lucide-svelte/icons/computer';
	import Menu from 'lucide-svelte/icons/menu';
	import * as Sheet from '../ui/sheet';
	import { Button } from '../ui/button';
	import { Input } from '../ui/input';
	import * as DropdownMenu from '../ui/dropdown-menu';
	import * as m from '$lib/paraglide/messages.js';
	import { cn } from '$lib/utils';
	import { page } from '$app/state';
	import { deLocalizeHref, setLocale } from '$lib/paraglide/runtime';
	import { useDebounce } from '$lib/hooks/useDebounce.svelte';
	import { getAppState } from '$lib/states/app.svelte';
	import { getUserState } from '$lib/states/user.svelte';

	interface Props {
		className?: string;
	}

	let { className }: Props = $props();

	let canonicalPath = $derived(deLocalizeHref(page.url.pathname));

	let search = $state('');
	let debouncedSearch = useDebounce(() => search, 500);

	const appState = getAppState();
	const userState = getUserState();
</script>

<header class="bg-background sticky top-0 z-10 flex h-16 items-center gap-4 border-b px-4 md:px-6">
	<nav
		class="hidden flex-col gap-6 text-lg font-medium md:flex md:flex-row md:items-center md:gap-5 md:text-sm lg:gap-6"
	>
		<a href={AppRoute.DASHBOARD} class="flex items-center gap-2 text-lg font-semibold md:text-base">
			<Paperclip class="h-6 w-6" />
			<span class="sr-only">DocPort</span>
		</a>
		<a
			href={AppRoute.DASHBOARD}
			class={cn(
				'text-muted-foreground hover:text-foreground transition-colors',
				canonicalPath === AppRoute.DASHBOARD && 'text-foreground'
			)}>{m.dashboard()}</a
		>
		<a
			href={AppRoute.PROJECTS}
			class={cn(
				'text-muted-foreground hover:text-foreground transition-colors',
				canonicalPath === AppRoute.PROJECTS && 'text-foreground'
			)}>{m.projects()}</a
		>
	</nav>
	<Sheet.Root>
		<Sheet.Trigger asChild let:builder>
			<Button variant="outline" size="icon" class="shrink-0 md:hidden" builders={[builder]}>
				<Menu class="h-5 w-5" />
				<span class="sr-only">{m.toggle_navigation_menu()}</span>
			</Button>
		</Sheet.Trigger>
		<Sheet.Content side="left">
			<nav class="grid gap-6 text-lg font-medium">
				<a href={AppRoute.DASHBOARD} class="flex items-center gap-2 text-lg font-semibold">
					<Paperclip class="h-6 w-6" />
					<span class="sr-only">DocPort</span>
				</a>
				<a
					href={AppRoute.DASHBOARD}
					class={cn(
						'text-muted-foreground hover:text-foreground',
						canonicalPath === AppRoute.DASHBOARD && 'text-foreground'
					)}>{m.dashboard()}</a
				>
				<a
					href={AppRoute.PROJECTS}
					class={cn(
						'text-muted-foreground hover:text-foreground',
						canonicalPath === AppRoute.PROJECTS && 'text-foreground'
					)}>{m.projects()}</a
				>
			</nav>
		</Sheet.Content>
	</Sheet.Root>
	<div class="flex w-full items-center gap-4 md:ml-auto md:gap-2 lg:gap-4">
		<form class="ml-auto flex-1 sm:flex-initial">
			<div class="relative">
				<Search class="text-muted-foreground absolute top-2.5 left-2.5 h-4 w-4" />
				<Input
					type="search"
					placeholder={m.search()}
					class="pl-8 sm:w-[300px] md:w-[200px] lg:w-[300px]"
					bind:value={search}
				/>
			</div>
		</form>
		<div>{m.fair_known_reindeer_bless({ name: userState.name })}</div>
		<DropdownMenu.Root>
			<DropdownMenu.Trigger asChild let:builder>
				<Button builders={[builder]} variant="secondary" size="icon" class="rounded-full">
					<CircleUser class="h-5 w-5" />
					<span class="sr-only">{m.toggle_user_menu()}</span>
				</Button>
			</DropdownMenu.Trigger>
			<DropdownMenu.Content align="end">
				<DropdownMenu.Label>{m.my_account()}</DropdownMenu.Label>
				<DropdownMenu.Separator />
				<DropdownMenu.Item>{m.settings()}</DropdownMenu.Item>
				<DropdownMenu.Item>{m.support()}</DropdownMenu.Item>
				<DropdownMenu.Separator />
				<DropdownMenu.Item>{m.logout()}</DropdownMenu.Item>
				<DropdownMenu.Separator />
				<DropdownMenu.Label>{m.language()}</DropdownMenu.Label>
				<DropdownMenu.Item onclick={() => setLocale('nl')}>{m.dutch()}</DropdownMenu.Item>
				<DropdownMenu.Item onclick={() => setLocale('en')}>{m.english()}</DropdownMenu.Item>
				<DropdownMenu.Separator />
				<DropdownMenu.Label>{m.theme()}</DropdownMenu.Label>
				<DropdownMenu.Item on:click={() => appState.activateLightTheme()}>
					<Sun class="mr-2 h-4 w-4" />
					<span>{m.light()}</span>
				</DropdownMenu.Item>
				<DropdownMenu.Item on:click={() => appState.activateDarkTheme()}>
					<Moon class="mr-2 h-4 w-4" />
					<span>{m.dark()}</span>
				</DropdownMenu.Item>
				<DropdownMenu.Item on:click={() => appState.activateSystemTheme()}>
					<Computer class="mr-2 h-4 w-4" />
					<span>{m.system()}</span>
				</DropdownMenu.Item>
			</DropdownMenu.Content>
		</DropdownMenu.Root>
	</div>
</header>
