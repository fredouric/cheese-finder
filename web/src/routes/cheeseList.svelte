<script lang="ts">
	import CheeseCard from './cheeseCard.svelte';
	import type { Cheese } from '../../gen/cheese/v1/cheese';

	export let cheeses: Cheese[];
</script>

<div class="scrollable-list">
	{#each cheeses as cheese}
		<CheeseCard {cheese} />
	{/each}
</div>
