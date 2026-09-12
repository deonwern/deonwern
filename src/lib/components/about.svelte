<script>
	import { resolve } from '$app/paths';
</script>

<section class="mx-auto max-w-3xl px-4 py-16 text-left">
	<div class="card border border-base-200 bg-base-100 shadow-sm">
		<div class="card-body p-8">
			<h2 class="mb-4 card-title text-2xl font-bold tracking-tight text-base-content">
				About the Journey
			</h2>
			<p class="mb-4 text-base leading-relaxed text-base-content/80">
				I am a solo developer who decided to bypass traditional corporate hiring and engineer a
				sustainable career from scratch. By building open-core, privacy-focused applications and
				documenting every step in public, I am proving that independent builders can create their
				own path.
			</p>
			<p class="mb-6 text-base leading-relaxed text-base-content/80">
				Every project I build is backed by a strict set of principles to ensure user ownership,
				transparency, and longevity.
			</p>
			<div class="card-actions">
				<a href={resolve('/oath')} class="btn btn-primary btn-sm"> Read My Unbreakable Oath → </a>
			</div>
		</div>
	</div>
</section>
