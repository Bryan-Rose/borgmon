
import type { ServerInit } from '@sveltejs/kit';
import { setContext } from 'svelte';
import { GetApiClient } from '$lib/api/api'

export const init: ServerInit = () => {
    // setContext('APICLIENT', GetApiClient());
};