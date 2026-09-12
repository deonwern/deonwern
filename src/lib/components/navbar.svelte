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

<header class="navbar mx-auto max-w-6xl border-b border-base-200 bg-base-100 px-6 py-6">
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

	<div class="navbar-center">
		<!-- Megamenu container -->
		<div class="megamenu hidden p-2 max-sm:megamenu-vertical md:flex" id="main-megamenu" popover>
			<span class="megamenu-active"></span>

			{#each menuItems as item, index (index)}
				<button popovertarget={item.id} class="cursor-pointer px-4 py-2 font-medium select-none">
					{item.label}
				</button>

				<div id={item.id} popover>
					<ul class="menu w-52 rounded-box border border-base-200 bg-base-100 p-2 shadow-lg">
						{#each item.children as sub, index (index)}
							<li><a href={resolve(sub.href)} class="rounded-btn px-4 py-2">{sub.name}</a></li>
						{/each}
					</ul>
				</div>
			{/each}
		</div>
	</div>

	<div class="navbar-end">
		<a href={resolve('/#connect')} class="btn btn-md btn-primary">Send Message</a>
	</div>

	<Fab />
</header>
