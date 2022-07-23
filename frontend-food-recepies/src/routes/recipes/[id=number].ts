/** @type {import('@sveltejs/kit').Load} */
export async function GET(page) {

    const apiKey = import.meta.env.VITE_SPOONACULAR_API_KEY;

    const url =
        `https://api.spoonacular.com/recipes/${page.params.id}/information?&apiKey=${apiKey}`;
    const res = await fetch(url);
    const data = await res.json();

    if (res.ok) {
        return {
            body: {
                recipe: data
            }
        }
    };

    return {
        body: {
            error: "невозможно получить рецепт"
        }
    }
}