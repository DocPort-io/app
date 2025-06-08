<script lang="ts">
	import { Menu } from '@lucide/svelte';
	import { page } from '$app/state';
	import DocPortLogo from '$lib/assets/logo.svg';
	import { Button } from '$lib/components/ui/button';
	import * as Sheet from '$lib/components/ui/sheet';
	import { AppRoute } from '$lib/constants';
	import { m } from '$lib/paraglide/messages';
	import { deLocalizeHref } from '$lib/paraglide/runtime';
	import { cn } from '$lib/utils';

	import ProfileMenu from './navigation-bar/profile-menu.svelte';
	import Search from './navigation-bar/search.svelte';
	import TeamSwitcher from './navigation-bar/team-switcher.svelte';

	let canonicalPath = $derived(deLocalizeHref(page.url.pathname));
</script>

<header class="bg-background sticky top-0 z-10 flex h-16 items-center gap-4 border-b px-4 md:px-6">
	<nav
		class="hidden flex-col gap-6 text-lg font-medium md:flex md:flex-row md:items-center md:gap-5 md:text-sm lg:gap-6"
	>
		<a
			href={AppRoute.DASHBOARD()}
			class="flex h-10 w-10 items-center gap-2 text-lg font-semibold md:text-base"
		>
			<img
				src={DocPortLogo}
				alt="DocPort Logo"
				class="aspect-square rounded-md object-cover"
				height="64"
				width="64"
			/>
			<span class="sr-only">DocPort</span>
		</a>
		<a
			href={AppRoute.DASHBOARD()}
			class={cn(
				'text-muted-foreground hover:text-foreground transition-colors',
				canonicalPath === AppRoute.DASHBOARD() && 'text-foreground'
			)}>{m.dashboard()}</a
		>
		<a
			href={AppRoute.PROJECTS()}
			class={cn(
				'text-muted-foreground hover:text-foreground transition-colors',
				canonicalPath === AppRoute.PROJECTS() && 'text-foreground'
			)}>{m.projects()}</a
		>
	</nav>
	<Sheet.Root>
		<Sheet.Trigger>
			<Button variant="outline" size="icon" class="shrink-0 md:hidden">
				<Menu class="h-5 w-5" />
				<span class="sr-only">{m.toggle_navigation_menu()}</span>
			</Button>
		</Sheet.Trigger>
		<Sheet.Content side="left">
			<Sheet.Header>
				<Sheet.Title>DocPort</Sheet.Title>
			</Sheet.Header>
			<nav class="grid gap-6 overflow-y-auto px-4 text-sm font-medium">
				<a
					href={AppRoute.DASHBOARD()}
					class={cn(
						'text-muted-foreground hover:text-foreground',
						canonicalPath === AppRoute.DASHBOARD() && 'text-foreground'
					)}>{m.dashboard()}</a
				>
				<a
					href={AppRoute.PROJECTS()}
					class={cn(
						'text-muted-foreground hover:text-foreground',
						canonicalPath === AppRoute.PROJECTS() && 'text-foreground'
					)}>{m.projects()}</a
				>
			</nav>
		</Sheet.Content>
	</Sheet.Root>
	<div class="flex w-full items-center gap-4 md:ml-auto md:gap-2 lg:gap-4">
		<Search />
		<TeamSwitcher />
		<ProfileMenu />
	</div>
</header>
