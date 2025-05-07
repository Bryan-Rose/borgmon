<script lang="ts">
	import { goto } from '$app/navigation';
	import { GetApiClient } from '$lib/api/api';
	import { onMount } from 'svelte';

	let email = '';
	let password = '';
	let errorMessage = '';
	let loading = false;

	async function handleSubmit() {
		loading = true;
		errorMessage = '';

		try {
			var client = GetApiClient();
			if (await client.Login(email, password, true)) {
				await goto('/'); // Redirect to dashboard
			} else {
				errorMessage = 'Invalid email or password.';
			}
		} catch (error) {
			errorMessage = 'An error occurred during login.';
			console.error('Login error:', error);
		} finally {
			loading = false;
		}
	}

	onMount(async () => {
		var client = GetApiClient();
		const valid = await client.TryUseValidToken();

		if (valid) {
			goto('/', { replaceState: true }); // Redirect if logged in
		}
	});
</script>

<div class="flex min-h-screen items-center justify-center bg-gray-100">
	<div class="mx-4 mt-4 bg-white px-8 py-6 text-left shadow-lg sm:w-3/4 md:w-1/3 lg:w-1/4">
		<h3 class="text-center text-2xl font-bold text-gray-800">Log in to your account</h3>
		<form class="mt-4" on:submit|preventDefault={handleSubmit}>
			<div>
				<label class="mb-2 block text-sm font-bold text-gray-700" for="email">Email</label>
				<input
					class="focus:shadow-outline w-full appearance-none rounded border px-3 py-2 leading-tight text-gray-700 shadow focus:outline-none"
					id="email"
					type="email"
					placeholder="Email"
					bind:value={email}
					required
				/>
			</div>
			<div class="mt-4">
				<label class="mb-2 block text-sm font-bold text-gray-700" for="password">Password</label>
				<input
					class="focus:shadow-outline w-full appearance-none rounded border px-3 py-2 leading-tight text-gray-700 shadow focus:outline-none"
					id="password"
					type="password"
					placeholder="Password"
					bind:value={password}
					required
				/>
			</div>
			{#if errorMessage}
				<p class="mt-4 text-xs italic text-red-500">{errorMessage}</p>
			{/if}
			<div class="mt-6">
				<button
					type="submit"
					class="w-full transform rounded-md bg-indigo-500 px-4 py-2 tracking-wide text-white transition-colors duration-200 hover:bg-indigo-400 focus:bg-indigo-400 focus:outline-none focus:ring focus:ring-indigo-300 disabled:cursor-not-allowed disabled:bg-gray-400"
					disabled={loading}
				>
					{#if loading}
						Logging in...
					{:else}
						Log in
					{/if}
				</button>
			</div>
			<div class="mt-4 text-center">
				<a
					class="inline-block align-baseline text-sm text-blue-500 hover:text-blue-800"
					href="/forgot-password"
				>
					Forgot Password?
				</a>
			</div>
		</form>
	</div>
</div>
