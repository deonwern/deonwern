<script>
	import { resolve } from '$app/paths';
	import Fab from './fab.svelte';
	const github = 'https://github.com/deonwern';

	const menuItems = [
		{
			id: 'menu-apps',
			label: 'Apps',
			children: [
				{ name: 'All Apps', href: '/apps' },
				{ name: 'Carnelian', href: '/apps/carnelian' }
			]
		},
		{
			id: 'menu-devlogs',
			label: 'Dev Logs',
			children: [
				{ name: 'All Dev Logs', href: '/blog' },
				{ name: 'Top', href: '/blog?top=true' }
			]
		},
		{
			id: 'menu-media',
			label: 'Media & Links',
			children: [
				{ name: 'Links', href: '/links' },
				{ name: 'YouTube', href: '/youtube' }
			]
		}
	];
</script>

<header class="navbar border-base-200 bg-base-100 mx-auto max-w-6xl border-b px-6 py-6">
	<div class="navbar-start">
		<a
			href={github}
			target="_blank"
			rel="noopener noreferrer"
			class="btn btn-ghost px-2 text-xl font-bold tracking-tight"
		>
			Deonwern
		</a>
	</div>

	<!-- Desktop & Mobile Megamenu Integration -->
	<div class="navbar-center">
		<!-- Mobile trigger button -->
		<button class="btn sm:hidden" popovertarget="main-megamenu">Menu</button>

		<!-- Megamenu container -->
		<div
			class="megamenu max-sm:megamenu-vertical border-base-300 hidden border p-2 md:flex"
			id="main-megamenu"
			popover
		>
			<span class="megamenu-active"></span>

			{#each menuItems as item}
				<button popovertarget={item.id} class="cursor-pointer px-4 py-2 font-medium select-none">
					{item.label}
				</button>

				<div id={item.id} popover>
					<ul class="menu rounded-box border-base-200 bg-base-100 w-52 border p-2 shadow-lg">
						{#each item.children as sub}
							<li><a href={resolve(sub.href)} class="rounded-btn px-4 py-2">{sub.name}</a></li>
						{/each}
					</ul>
				</div>
			{/each}
		</div>
	</div>

	<div class="navbar-end">
		<a href={resolve('/#connect')} class="btn btn-outline btn-md">Connect</a>
	</div>

	<Fab />
</header>
