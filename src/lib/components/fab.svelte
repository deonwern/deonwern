<script lang="ts">
	import { page } from '$app/state';
	import { resolve } from '$app/paths';
	import { browser } from '$app/environment';

	// Track scroll position
	let scrollY = $state(0);

	// Determine if we are currently on the homepage
	let isHome = $derived(page.url.pathname === '/');

	// Determine if the button should be visible
	// Visible if we've scrolled down OR if we are on a non-home page
	let showFab = $derived((browser && scrollY > 300) || !isHome);

	function handleClick(e: Event) {
		if (!isHome) {
			// If not home, let the link handle navigation back to home "/"
			return;
		}
		// If on home and scrolled down, scroll smoothly back to top
		e.preventDefault();
		window.scrollTo({ top: 0, behavior: 'smooth' });
	}
</script>

<svelte:window bind:scrollY />

{#if showFab}
	<div class="animate-fade-in fixed right-6 bottom-6 z-50 transition-all duration-300">
		{#if isHome}
			<button
				onclick={handleClick}
				class="btn rounded-full px-5 shadow-lg btn-neutral btn-sm md:btn-md"
				aria-label="Back to top"
			>
				Back to Top
			</button>
		{:else}
			<a
				href={resolve('/')}
				class="btn rounded-full px-5 shadow-lg btn-neutral btn-sm md:btn-md"
				aria-label="Back to home"
			>
				Back to Home
			</a>
		{/if}
	</div>
{/if}
