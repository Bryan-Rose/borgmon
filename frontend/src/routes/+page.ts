import type { RepoRecord } from '$lib/ repo-types';
import type { CollectionsListResponse } from '$lib/pb-types';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch }) => {
	const res = await fetch(`http://127.0.0.1:8090/api/collections/repos/records`, {
		headers: {
			"Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2xsZWN0aW9uSWQiOiJwYmNfMzE0MjYzNTgyMyIsImV4cCI6MTc0NDIxODEyNCwiaWQiOiI0NWN4YnV1MDk1N3dkd3AiLCJyZWZyZXNoYWJsZSI6dHJ1ZSwidHlwZSI6ImF1dGgifQ.pjDYOqw1kdbijqubiBBqNLXElraSaoQA-kfoPJKcxOk"
		}
	});
	
	const items = await res.json() as CollectionsListResponse<RepoRecord>;

	return {
		repos: items.items
	};
};