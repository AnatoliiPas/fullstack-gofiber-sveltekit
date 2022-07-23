<script>
	import '../app.css';
	import Nav from '../components/Nav.svelte';
	import Footer from '../components/Footer.svelte';
</script>

<Nav />

<main class="flex min-h-[52.5em]">
	<div class="m-auto">
		<slot />
	</div>
</main>

<Footer />
