<script lang="ts">
	import type { Cheese } from '../../gen/cheese/v1/cheese';

	export let cheese: Cheese;
</script>

<a href="/{cheese.id}">
	<div class="group mx-auto bg-slate-100 hover:bg-slate-200 rounded-xl p-4">
		<p>{cheese.fromage}</p>
		<p class="">{cheese.departement}</p>
	</div>
</a>
