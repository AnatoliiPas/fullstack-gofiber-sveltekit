<script>
	export let recipes;
</script>

<head>
	<title>Рецепты</title>
</head>
<div>
	<h1 class="text-2xl text-center mb-10 mt-5">Рецепты</h1>
	<div class="grid grid-cols-3 gap-10 m-5">
		{#each recipes as item}
			<a class="relative" href={`/recipesv2/${item.id}`} sveltekit:reload>
				<img
					class="shadow-lg w-32 h-32 rounded-sm md:w-80 md:h-80 md:rounded-xl"
					src={`http://127.0.0.1:8000/img/${item.image}`}
					alt={item.name}
				/>
				<p
					class="md:text-xl text-sm absolute z-20 top-0 mt-5 md:mt-20 text-center inset-x-0 text-white"
				>
					{item.name}
				</p>
			</a>
		{/each}
	</div>
</div>
