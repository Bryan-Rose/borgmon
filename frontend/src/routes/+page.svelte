<script lang="ts">
	import type { PageProps } from './$types';
	import { formatDate, formatBytes } from '$lib/util';

	let { data }: PageProps = $props();
</script>

{#each data.repos as repo}
	<div class="repo-record">
		<h2>{repo.name}</h2>
		<p>Updated: {formatDate(repo.borgUpdated)}</p>
		<p><strong>Path:</strong> {repo.path}</p>
		<p><strong>Last Backup:</strong> {formatDate(repo.lastBackup)}</p>

		{#if repo.borgInfo.cache}
			<h5>Cache Stats</h5>
			<p><strong>Total Chunks:</strong> {repo.borgInfo.cache.stats?.total_chunks}</p>
			<p>
				<strong>Total Compressed Size:</strong>
				{formatBytes(repo.borgInfo.cache.stats?.total_csize)}
			</p>
			<p><strong>Total Size:</strong> {formatBytes(repo.borgInfo.cache.stats?.total_size)}</p>
			<p>
				<strong>Total Unique Chunks:</strong>
				{repo.borgInfo.cache.stats?.total_unique_chunks}
			</p>
			<p>
				<strong>Unique Compressed Size:</strong>
				{formatBytes(repo.borgInfo.cache.stats?.unique_csize)}
			</p>
			<p><strong>Unique Size:</strong> {formatBytes(repo.borgInfo.cache.stats?.unique_size)}</p>
		{:else}
			<p>Cache information not available.</p>
		{/if}

		
		<h3>Borg List</h3>
		{#if repo.borgList}
			<h4>Archives</h4>
			{#if repo.borgList.archives && repo.borgList.archives.length > 0}
				<ul>
					{#each repo.borgList.archives as archive}
						<li>
							<strong>Archive:</strong>
							{archive.archive} ({archive.name})<br />
							<!-- <strong>ID:</strong>
							{archive.id}<br /> -->
							<strong>Start Time:</strong>
							{formatDate(archive.start)}<br />
							<strong>Time:</strong>
							{formatDate(archive.time)}<br />
							<strong>Barchive:</strong>
							{archive.barchive}
						</li>
					{/each}
				</ul>
			{:else}
				<p>No archives found.</p>
			{/if}
		{:else}
			<p>Borg list information not available.</p>
		{/if}
	</div>
{/each}

<style>
	.repo-record {
		border: 1px solid #ccc;
		padding: 15px;
		margin-bottom: 15px;
		border-radius: 5px;
		background-color: #f9f9f9;
	}

	.repo-record h2 {
		margin-top: 0;
		margin-bottom: 10px;
		color: #333;
	}

	.repo-record h3 {
		margin-top: 20px;
		margin-bottom: 10px;
		color: #555;
	}

	.repo-record h4 {
		margin-top: 15px;
		margin-bottom: 5px;
		color: #777;
	}

	.repo-record p {
		margin-bottom: 5px;
	}

	.repo-record ul {
		list-style-type: disc;
		padding-left: 20px;
	}

	.repo-record li {
		margin-bottom: 8px;
		padding: 8px;
		border: 1px solid #eee;
		border-radius: 3px;
		background-color: #fff;
	}

	.repo-record strong {
		font-weight: bold;
	}
</style>
