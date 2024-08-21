<script lang="ts">
	import type { Cheese } from '../../gen/cheese/v1/cheese';

	export let cheese: Cheese;
</script>

<div class="cheese-card">
	<h3>{cheese.fromage}</h3>
	<p>Type: {cheese.departement}</p>
</div>
