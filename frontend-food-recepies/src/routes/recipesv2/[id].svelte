<script>
	export let recipe;
</script>

<head>
	<title>Рецепт</title>
</head>

<h1 class="text-2xl text-center mb-10 mt-5">{recipe.name}</h1>
<div class="md:flex md:justify-between md:mx-20 my-10">
	<img
		class="mx-auto rounded-xl shadow-xl w-40 h-40 md:w-[400px] md:h-[400px]"
		src={`http://127.0.0.1:8000/img/${recipe.image}`}
		alt={recipe.name}
	/>
	<div class="mx-10">
		<p>{recipe.descriptions}</p>
		<div class="my-5">
			<h2 class="text-xl font-bold">Индигриенты:</h2>
			{#each recipe.products as product}
				<p class="text-gray-600 text-sm">{product}</p>
			{/each}
		</div>
		<div class="my-5">
			<h2 class="text-xl font-bold">Пошаговое руководство:</h2>
			{#each recipe.manual as manual}
				<p class="text-gray-600 text-sm md:text-lg">{manual}</p>
			{/each}
		</div>
	</div>
</div>
