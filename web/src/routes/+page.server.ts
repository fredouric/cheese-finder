import { cheeseClient } from '$lib/server/api'
import type { PageServerLoad } from './$types';


export const load = (async ({ }) => {

	const { response } = await cheeseClient.getOneCheese({
		id: "10",
	});
	if (!response.cheese) {
		throw new Error('no data');
	}
	return {
        id: "10",
	};
}) satisfies PageServerLoad;
