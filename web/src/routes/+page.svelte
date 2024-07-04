<h1>Welcome to CheeseFinder</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>

<script lang="ts">
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<svelte:head>
	<title>Cheese {data.id}</title>
</svelte:head>

<p>Cheese ID: {data.id}</p>

