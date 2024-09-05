<script lang="ts">
	import type { PageData } from './$types';
	import CheeseList from './cheeseList.svelte';

	export let data: PageData;
</script>

<h1 class="text-3xl font-bold-underline">Welcome to CheeseFinder</h1>

<svelte:head>
	<title>Cheese List</title>
</svelte:head>

<CheeseList cheeses={data.cheeses} />
