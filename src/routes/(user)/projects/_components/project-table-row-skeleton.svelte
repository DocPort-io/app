<script lang="ts">
	import { Skeleton } from '$lib/components/ui/skeleton';
	import * as Table from '$lib/components/ui/table';
</script>

<Table.Row>
	<Table.Cell class="font-medium">
		<Skeleton class="h-4 w-[200px] md:w-[300px]" />
	</Table.Cell>
	<Table.Cell class="hidden md:table-cell"><Skeleton class="h-4 w-[200px]" /></Table.Cell>
	<Table.Cell></Table.Cell>
</Table.Row>
