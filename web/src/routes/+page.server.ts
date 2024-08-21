import { cheeseClient } from '$lib/server/api'
import type { PageServerLoad } from './$types';
import clone from 'just-clone';


export const load = (async ({ }) => {

    const { response } = await cheeseClient.getAllCheeses({
        offset: "10",
        limit: "10",
    });
    if (!response.data) {
        throw new Error('no data');
    }

    return {
        cheeses: clone(response.data),
    };
}) satisfies PageServerLoad;
