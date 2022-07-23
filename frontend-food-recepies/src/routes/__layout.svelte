<script>
	import '../app.css';
	import Nav from '../components/Nav.svelte';
	import Footer from '../components/Footer.svelte';
</script>

<div class="flex flex-col h-screen">
	<Nav />

	<main class="flex flex-grow">
		<div class="m-auto">
			<slot />
		</div>
	</main>

	<Footer />
</div>
