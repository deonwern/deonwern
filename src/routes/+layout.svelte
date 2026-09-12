<script lang="ts">
	import './layout.css';
	import favicon from '$lib/assets/favicon.svg';
	import { Navbar, Footer } from '$lib/index';

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
	<title>deonwern portfolio</title>
</svelte:head>

<Navbar />
{@render children()}
<Footer />
