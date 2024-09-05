<script lang="ts">
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<div>
	<h1>{data.cheese.id}</h1>
	<h1>{data.cheese.fromage}</h1>
	<h1>{data.cheese.lait}</h1>
	<h1>{data.cheese.departement}</h1>
	<h1>{data.cheese.geoPoint2D}</h1>
</div>
