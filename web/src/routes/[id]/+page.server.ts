import { cheeseClient } from '$lib/server/api'
import clone from 'just-clone';
import type { PageServerLoad } from './$types';


export const load = (async ({ params }) => {
    const { response } = await cheeseClient.getOneCheese({
        id: params.id,
    });
    if (!response.cheese) {
        throw new Error('no data');
    }

    return {
        cheese: clone(response.cheese),
    };
}) satisfies PageServerLoad
