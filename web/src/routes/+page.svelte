<script lang="ts">
	import type { PageData } from './$types';
	import CheeseList from './cheeseList.svelte';

	export let data: PageData;
</script>

<h1>Welcome to CheeseFinder</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>

<svelte:head>
	<title>Cheese List</title>
</svelte:head>

<CheeseList cheeses={data.cheeses} />
