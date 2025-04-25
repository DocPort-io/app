<script lang="ts">
	import { Computer, Moon, Sun } from '@lucide/svelte';
	import * as Avatar from '$lib/components/ui/avatar';
	import { Button } from '$lib/components/ui/button';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import { m } from '$lib/paraglide/messages';
	import { setLocale } from '$lib/paraglide/runtime';
	import { getUserState } from '$lib/stores/user.svelte';
	import { setMode } from 'mode-watcher';

	const userState = getUserState();

	let initials = $derived.by(() => {
		const splitName = userState.name.toUpperCase().split(' ');
		if (splitName.length === 1) return splitName[0][0];
		return `${splitName[0][0]}${splitName[splitName.length - 1][0]}`;
	});
</script>

<DropdownMenu.Root>
	<DropdownMenu.Trigger>
		<Button variant="secondary" size="icon" class="rounded-full">
			<Avatar.Root>
				<Avatar.Image src={userState.avatarUrl} alt={initials} />
				<Avatar.Fallback>{initials}</Avatar.Fallback>
			</Avatar.Root>
			<span class="sr-only">{m.toggle_user_menu()}</span>
		</Button>
	</DropdownMenu.Trigger>
	<DropdownMenu.Content align="end">
		<DropdownMenu.Label>{m.my_account()}</DropdownMenu.Label>
		<DropdownMenu.Separator />
		<DropdownMenu.Item>{m.settings()}</DropdownMenu.Item>
		<DropdownMenu.Item>{m.support()}</DropdownMenu.Item>
		<DropdownMenu.Separator />
		<DropdownMenu.Item onclick={() => userState.logout()}>{m.logout()}</DropdownMenu.Item>
		<DropdownMenu.Separator />
		<DropdownMenu.Label>{m.language()}</DropdownMenu.Label>
		<DropdownMenu.Item onclick={() => setLocale('nl')}>{m.dutch()}</DropdownMenu.Item>
		<DropdownMenu.Item onclick={() => setLocale('en')}>{m.english()}</DropdownMenu.Item>
		<DropdownMenu.Separator />
		<DropdownMenu.Label>{m.theme()}</DropdownMenu.Label>
		<DropdownMenu.Item onclick={() => setMode('light')}>
			<Sun class="mr-2 h-4 w-4" />
			<span>{m.light()}</span>
		</DropdownMenu.Item>
		<DropdownMenu.Item onclick={() => setMode('dark')}>
			<Moon class="mr-2 h-4 w-4" />
			<span>{m.dark()}</span>
		</DropdownMenu.Item>
		<DropdownMenu.Item onclick={() => setMode('system')}>
			<Computer class="mr-2 h-4 w-4" />
			<span>{m.system()}</span>
		</DropdownMenu.Item>
	</DropdownMenu.Content>
</DropdownMenu.Root>
