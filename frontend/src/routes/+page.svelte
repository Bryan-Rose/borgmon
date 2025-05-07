<script lang="ts">
	import type { PageProps } from './$types';
	import { formatDate, formatBytes } from '$lib/util';

	let { data }: PageProps = $props();
</script>

<div class="grid grid-cols-1 gap-4 md:grid-cols-2">
	{#each data.repos as repo (repo.path)}
		<div class="rounded-lg bg-white p-6 shadow-md">
			<h2 class="mb-2 text-xl font-semibold text-gray-800">{repo.name}</h2>
			<p class="mb-1 text-gray-600">
				<strong class="font-medium">Path:</strong>
				{repo.path}
			</p>
			<p class="mb-1 text-gray-600">
				<strong class="font-medium">Borg Updated:</strong>
				{formatDate(repo.borgUpdated)}
			</p>
			<p class="mb-1 text-gray-600">
				<strong class="font-medium">Last Backup:</strong>
				{formatDate(repo.lastBackup)}
			</p>
		</div>

		<div class="rounded-lg bg-gray-100 p-6 shadow-md">
			<h3 class="mb-2 text-lg font-semibold text-gray-700">Detailed Information</h3>

			<div class="mb-4">
				{#if repo.borgInfo}
					<div class="mb-2">
						<h5 class="text-sm font-semibold text-gray-700">Repository</h5>
						<p class="text-gray-600">
							<strong class="font-medium">Location:</strong>
							{repo.borgInfo.repository?.location || '-'}
						</p>
						<p class="text-gray-600">
							<strong class="font-medium">Last Modified:</strong>
							{formatDate(repo.borgInfo.repository?.last_modified)}
						</p>
					</div>

					<div>
						{#if repo.borgInfo.cache}
							<h6 class="mt-1 text-xs font-semibold text-gray-700">Stats</h6>
							<div class="flex items-center space-x-4 text-sm">
								<p class="text-gray-600">
									<strong class="font-medium">Total Chunks:</strong>
									{repo.borgInfo.cache.stats?.total_chunks || '-'}
								</p>
								<p class="text-gray-600">
									<strong class="font-medium">Total Size:</strong>
									{formatBytes(repo.borgInfo.cache.stats?.total_size)}
								</p>
								<p class="text-gray-600">
									<strong class="font-medium">Total CSize:</strong>
									{formatBytes(repo.borgInfo.cache.stats?.total_csize)}
								</p>
							</div>
							<div class="mt-1 flex items-center space-x-4 text-sm">
								<p class="text-gray-600">
									<strong class="font-medium">Unique Chunks:</strong>
									{repo.borgInfo.cache.stats?.total_unique_chunks || '-'}
								</p>
								<p class="text-gray-600">
									<strong class="font-medium">Unique Size:</strong>
									{formatBytes(repo.borgInfo.cache.stats.unique_size)}
								</p>
								<p class="text-gray-600">
									<strong class="font-medium">Unique CSize:</strong>
									{formatBytes(repo.borgInfo.cache.stats?.unique_csize)}
								</p>
							</div>
						{:else}
							<p class="text-gray-500">Cache information not available.</p>
						{/if}
					</div>
				{:else}
					<p class="text-gray-500">Borg info not available.</p>
				{/if}
			</div>

			<div class="mb-4">
				<h4 class="text-md mb-1 font-semibold text-gray-700">Archives</h4>
				{#if repo.borgList?.archives?.length > 0}
					<ul class="list-disc pl-5 text-gray-600">
						{#each repo.borgList.archives as archive}
							<li>
								<strong class="font-medium">{archive.name}</strong>
								<div class="ml-4 flex items-center space-x-2 text-sm">
									<span class="font-semibold">Start:</span>
									<span>{formatDate(archive.start)} </span>
									<span class="font-semibold">End:</span>
									<span>{formatDate(archive.time)}</span>
								</div>
							</li>
						{/each}
					</ul>
				{:else}
					<p class="text-gray-500">No archives found.</p>
				{/if}
			</div>
		</div>
	{/each}
</div>
