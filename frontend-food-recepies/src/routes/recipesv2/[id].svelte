<script>
	export let recipe;
</script>

<head>
	<title>Рецепт</title>
</head>

<div style="width: 80%; margin: 2em auto;">
	<h1>{recipe.name}</h1>
	<img src={`http://127.0.0.1:8000/img/${recipe.image}`} alt={recipe.name} />
	<p>{recipe.descriptions}</p>
	<br />
	<h2>Индигриенты:</h2>
	{#each recipe.products as product}
		<p>{product}</p>
	{/each}
	<br />
	<h2>Пошаговое руководство:</h2>
	{#each recipe.manual as manual}
		<p>{manual}</p>
	{/each}
</div>
