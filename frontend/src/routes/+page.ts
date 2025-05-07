import type { RepoRecord } from '$lib/api/ repo-types';
import { GetApiClient } from '$lib/api/api';
import type { CollectionsListResponse } from '$lib/api/pb-types';
import type { PageLoad } from './$types';

export const load: PageLoad = async () => {
	const apiClient = GetApiClient();
	const repos = await apiClient.ListAllRepos();
	return {
		repos: repos
	};
};