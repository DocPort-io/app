<script lang="ts">
	import { Activity } from '@lucide/svelte';
	import { ArrowUpRight } from '@lucide/svelte';
	import { CreditCard } from '@lucide/svelte';
	import { DollarSign } from '@lucide/svelte';
	import { Users } from '@lucide/svelte';
	import UserPageLayout from '$lib/components/layouts/user-page-layout.svelte';
	import * as Avatar from '$lib/components/ui/avatar';
	import { Badge } from '$lib/components/ui/badge';
	import { Button } from '$lib/components/ui/button';
	import * as Card from '$lib/components/ui/card';
	import * as Table from '$lib/components/ui/table';
</script>

<UserPageLayout title="Dashboard">
	<div class="grid gap-4 md:grid-cols-2 md:gap-8 lg:grid-cols-4">
		<Card.Root
			data-x-chunk-name="dashboard-01-chunk-0"
			data-x-chunk-description="A card showing the total revenue in USD and the percentage difference from last month."
		>
			<Card.Header class="flex flex-row items-center justify-between space-y-0 pb-2">
				<Card.Title class="text-sm font-medium">Total Revenue</Card.Title>
				<DollarSign class="text-muted-foreground h-4 w-4" />
			</Card.Header>
			<Card.Content>
				<div class="text-2xl font-bold">$45,231.89</div>
				<p class="text-muted-foreground text-xs">+20.1% from last month</p>
			</Card.Content>
		</Card.Root>
		<Card.Root
			data-x-chunk-name="dashboard-01-chunk-1"
			data-x-chunk-description="A card showing the total subscriptions and the percentage difference from last month."
		>
			<Card.Header class="flex flex-row items-center justify-between space-y-0 pb-2">
				<Card.Title class="text-sm font-medium">Subscriptions</Card.Title>
				<Users class="text-muted-foreground h-4 w-4" />
			</Card.Header>
			<Card.Content>
				<div class="text-2xl font-bold">+2350</div>
				<p class="text-muted-foreground text-xs">+180.1% from last month</p>
			</Card.Content>
		</Card.Root>
		<Card.Root
			data-x-chunk-name="dashboard-01-chunk-2"
			data-x-chunk-description="A card showing the total sales and the percentage difference from last month."
		>
			<Card.Header class="flex flex-row items-center justify-between space-y-0 pb-2">
				<Card.Title class="text-sm font-medium">Sales</Card.Title>
				<CreditCard class="text-muted-foreground h-4 w-4" />
			</Card.Header>
			<Card.Content>
				<div class="text-2xl font-bold">+12,234</div>
				<p class="text-muted-foreground text-xs">+19% from last month</p>
			</Card.Content>
		</Card.Root>
		<Card.Root
			data-x-chunk-name="dashboard-01-chunk-3"
			data-x-chunk-description="A card showing the total active users and the percentage difference from last hour."
		>
			<Card.Header class="flex flex-row items-center justify-between space-y-0 pb-2">
				<Card.Title class="text-sm font-medium">Active Now</Card.Title>
				<Activity class="text-muted-foreground h-4 w-4" />
			</Card.Header>
			<Card.Content>
				<div class="text-2xl font-bold">+573</div>
				<p class="text-muted-foreground text-xs">+201 since last hour</p>
			</Card.Content>
		</Card.Root>
	</div>
	<div class="grid gap-4 md:gap-8 lg:grid-cols-2 xl:grid-cols-3">
		<Card.Root
			class="xl:col-span-2"
			data-x-chunk-name="dashboard-01-chunk-4"
			data-x-chunk-description="A card showing a table of recent transactions with a link to view all transactions."
		>
			<Card.Header class="flex flex-row items-center">
				<div class="grid gap-2">
					<Card.Title>Transactions</Card.Title>
					<Card.Description>Recent transactions from your store.</Card.Description>
				</div>
				<Button href="##" size="sm" class="ml-auto gap-1">
					View All
					<ArrowUpRight class="h-4 w-4" />
				</Button>
			</Card.Header>
			<Card.Content>
				<Table.Root>
					<Table.Header>
						<Table.Row>
							<Table.Head>Customer</Table.Head>
							<Table.Head class="xl:table.-column hidden">Type</Table.Head>
							<Table.Head class="xl:table.-column hidden">Status</Table.Head>
							<Table.Head class="xl:table.-column hidden">Date</Table.Head>
							<Table.Head class="text-right">Amount</Table.Head>
						</Table.Row>
					</Table.Header>
					<Table.Body>
						<Table.Row>
							<Table.Cell>
								<div class="font-medium">Liam Johnson</div>
								<div class="text-muted-foreground hidden text-sm md:inline">liam@example.com</div>
							</Table.Cell>
							<Table.Cell class="xl:table.-column hidden">Sale</Table.Cell>
							<Table.Cell class="xl:table.-column hidden">
								<Badge class="text-xs" variant="outline">Approved</Badge>
							</Table.Cell>
							<Table.Cell class="md:table.-cell xl:table.-column hidden lg:hidden">
								2023-06-23
							</Table.Cell>
							<Table.Cell class="text-right">$250.00</Table.Cell>
						</Table.Row>
						<Table.Row>
							<Table.Cell>
								<div class="font-medium">Olivia Smith</div>
								<div class="text-muted-foreground hidden text-sm md:inline">olivia@example.com</div>
							</Table.Cell>
							<Table.Cell class="xl:table.-column hidden">Refund</Table.Cell>
							<Table.Cell class="xl:table.-column hidden">
								<Badge class="text-xs" variant="outline">Declined</Badge>
							</Table.Cell>
							<Table.Cell class="md:table.-cell xl:table.-column hidden lg:hidden">
								2023-06-24
							</Table.Cell>
							<Table.Cell class="text-right">$150.00</Table.Cell>
						</Table.Row>
						<Table.Row>
							<Table.Cell>
								<div class="font-medium">Noah Williams</div>
								<div class="text-muted-foreground hidden text-sm md:inline">noah@example.com</div>
							</Table.Cell>
							<Table.Cell class="xl:table.-column hidden">Subscription</Table.Cell>
							<Table.Cell class="xl:table.-column hidden">
								<Badge class="text-xs" variant="outline">Approved</Badge>
							</Table.Cell>
							<Table.Cell class="md:table.-cell xl:table.-column hidden lg:hidden">
								2023-06-25
							</Table.Cell>
							<Table.Cell class="text-right">$350.00</Table.Cell>
						</Table.Row>
						<Table.Row>
							<Table.Cell>
								<div class="font-medium">Emma Brown</div>
								<div class="text-muted-foreground hidden text-sm md:inline">emma@example.com</div>
							</Table.Cell>
							<Table.Cell class="xl:table.-column hidden">Sale</Table.Cell>
							<Table.Cell class="xl:table.-column hidden">
								<Badge class="text-xs" variant="outline">Approved</Badge>
							</Table.Cell>
							<Table.Cell class="md:table.-cell xl:table.-column hidden lg:hidden">
								2023-06-26
							</Table.Cell>
							<Table.Cell class="text-right">$450.00</Table.Cell>
						</Table.Row>
						<Table.Row>
							<Table.Cell>
								<div class="font-medium">Liam Johnson</div>
								<div class="text-muted-foreground hidden text-sm md:inline">liam@example.com</div>
							</Table.Cell>
							<Table.Cell class="xl:table.-column hidden">Sale</Table.Cell>
							<Table.Cell class="xl:table.-column hidden">
								<Badge class="text-xs" variant="outline">Approved</Badge>
							</Table.Cell>
							<Table.Cell class="md:table.-cell xl:table.-column hidden lg:hidden">
								2023-06-27
							</Table.Cell>
							<Table.Cell class="text-right">$550.00</Table.Cell>
						</Table.Row>
					</Table.Body>
				</Table.Root>
			</Card.Content>
		</Card.Root>
		<Card.Root
			data-x-chunk-name="dashboard-01-chunk-5"
			data-x-chunk-description="A card showing a list of recent sales with customer names and email addresses."
		>
			<Card.Header>
				<Card.Title>Recent Sales</Card.Title>
			</Card.Header>
			<Card.Content class="grid gap-8">
				<div class="flex items-center gap-4">
					<Avatar.Root class="hidden h-9 w-9 sm:flex">
						<Avatar.Image src="/avatars/01.png" alt="Avatar" />
						<Avatar.Fallback>OM</Avatar.Fallback>
					</Avatar.Root>
					<div class="grid gap-1">
						<p class="text-sm leading-none font-medium">Olivia Martin</p>
						<p class="text-muted-foreground text-sm">olivia.martin@email.com</p>
					</div>
					<div class="ml-auto font-medium">+$1,999.00</div>
				</div>
				<div class="flex items-center gap-4">
					<Avatar.Root class="hidden h-9 w-9 sm:flex">
						<Avatar.Image src="/avatars/02.png" alt="Avatar" />
						<Avatar.Fallback>JL</Avatar.Fallback>
					</Avatar.Root>
					<div class="grid gap-1">
						<p class="text-sm leading-none font-medium">Jackson Lee</p>
						<p class="text-muted-foreground text-sm">jackson.lee@email.com</p>
					</div>
					<div class="ml-auto font-medium">+$39.00</div>
				</div>
				<div class="flex items-center gap-4">
					<Avatar.Root class="hidden h-9 w-9 sm:flex">
						<Avatar.Image src="/avatars/03.png" alt="Avatar" />
						<Avatar.Fallback>IN</Avatar.Fallback>
					</Avatar.Root>
					<div class="grid gap-1">
						<p class="text-sm leading-none font-medium">Isabella Nguyen</p>
						<p class="text-muted-foreground text-sm">isabella.nguyen@email.com</p>
					</div>
					<div class="ml-auto font-medium">+$299.00</div>
				</div>
				<div class="flex items-center gap-4">
					<Avatar.Root class="hidden h-9 w-9 sm:flex">
						<Avatar.Image src="/avatars/04.png" alt="Avatar" />
						<Avatar.Fallback>WK</Avatar.Fallback>
					</Avatar.Root>
					<div class="grid gap-1">
						<p class="text-sm leading-none font-medium">William Kim</p>
						<p class="text-muted-foreground text-sm">will@email.com</p>
					</div>
					<div class="ml-auto font-medium">+$99.00</div>
				</div>
				<div class="flex items-center gap-4">
					<Avatar.Root class="hidden h-9 w-9 sm:flex">
						<Avatar.Image src="/avatars/05.png" alt="Avatar" />
						<Avatar.Fallback>SD</Avatar.Fallback>
					</Avatar.Root>
					<div class="grid gap-1">
						<p class="text-sm leading-none font-medium">Sofia Davis</p>
						<p class="text-muted-foreground text-sm">sofia.davis@email.com</p>
					</div>
					<div class="ml-auto font-medium">+$39.00</div>
				</div>
			</Card.Content>
		</Card.Root>
	</div>
</UserPageLayout>
