<script>
	export let recipe;
</script>

<head>
	<title>Рецепт</title>
</head>

<div class="text-center sm:text-left">
	<h1>{recipe.title}</h1>
	<img src={recipe.image} alt={recipe.title} />

	<div>
		<h2>Индигриенты</h2>
		{#each recipe.extendedIngredients as ingredients}
			<p>{ingredients.name}: {ingredients.amount}</p>
		{/each}
	</div>

	{#if recipe.cuisines.length > 0}
		<div>
			<h2>Cuisines</h2>
			{#each recipe.cuisines as cuisines}
				<p>{cuisines}</p>
			{/each}
		</div>
	{/if}

	{#if recipe.diets.length > 0}
		<div>
			<h2>Diets</h2>
			{#each recipe.diets as diets}
				<p>{diets}</p>
			{/each}
		</div>
	{/if}

	{#if recipe.dishTypes.length > 0}
		<div>
			<h2>DishTypes</h2>
			{#each recipe.dishTypes as dishTypes}
				<p>{dishTypes}</p>
			{/each}
		</div>
	{/if}

	{#if recipe.occasions.length > 0}
		<div>
			<h2>Occasions</h2>
			{#each recipe.occasions as occasions}
				<p>{occasions}</p>
			{/each}
		</div>
	{/if}
</div>
